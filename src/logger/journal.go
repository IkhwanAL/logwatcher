package logger

import (
	"bufio"
	"encoding/json"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Journal struct {
	MESSAGE              string
	PRIORITY             string
	REALTIME_TIMESTAMP   string `json:"__REALTIME_TIMESTAMP"`
	SYSTEMD_UNIT         string `json:"_SYSTEMD_UNIT"`
	COMM                 string `json:"_COMM"`
	TRANSPORT            string `json:"_TRANSPORT"`
	COREDUMP_SIGNAL_NAME string
	COREDUMP_COMM        string
	COREDUMP_PID         string
	BOOT_ID              string `json:"_BOOT_ID"`
	rawLine              string
}

func (j *Journal) getTimestamp() int64 {
	v, err := strconv.ParseInt(j.REALTIME_TIMESTAMP, 10, 64)
	if err != nil {
		log.Printf("Failed to convert string to int64 [REALTIME_TIMESTAMP]: %v e: %v", v, err.Error())
		return time.Now().UnixMicro()
	}

	return v
}

func (j *Journal) getDateTime() string {
	sec := j.getTimestamp()
	t := time.UnixMicro(sec)

	formattedValue := t.Format(time.DateTime)

	return formattedValue
}

func (j *Journal) getPriority() string {
	p, err := strconv.Atoi(j.PRIORITY)
	if err != nil {
		log.Printf("Failed to convert string to number [PRIORITY]: %v e: %v", j.PRIORITY, err.Error())
		p = 9999
	}

	switch p {
	case 1:
	case 2:
		return "[[red]CRITICAL[-]]"
	case 3:
		return "[[orange]ERROR[-]]"
	case 4:
		return "[[yellow]WARNING[-]]"
	case 5:
	case 6:
		return "[[white]INFO[-]]"
	case 7:
		return "[[gray]DEBUG[-]]"
	default:
		return "[[blue]NONE[-]]"
	}

	return ""
}

func (j *Journal) isCoreDump() bool {
	if strings.Contains(j.rawLine, "COREDUMP_") {
		return true
	}

	return false
}

func (j *Journal) constructCoreDumpMessage() string {
	var builder strings.Builder
	builder.WriteString("code dump: " + j.COREDUMP_COMM)
	builder.WriteString(" ( pid: " + j.COREDUMP_PID + ")" + " Look at \"journalctl -p crit\" ")
	builder.WriteString("signal: " + j.COREDUMP_SIGNAL_NAME)

	return builder.String()
}

func (j *Journal) getMessage() string {
	if j.isCoreDump() {
		return j.constructCoreDumpMessage()
	}

	return j.MESSAGE
}

func (j *Journal) Parse() string {
	var builder strings.Builder

	builder.WriteString("[" + j.getDateTime() + "] ")
	builder.WriteString(j.TRANSPORT + " ")
	builder.WriteString(j.getPriority() + " ")
	builder.WriteString(j.getMessage())

	return builder.String()
}

// Creating A Pipeline And Start Plumbing
func WatchJournal() <-chan Log {
	ch := make(chan Log)
	cmd := exec.Command("journalctl", "-f", "", "--output=json")
	out, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal("Failed To Pipe In To Log TUI ")
	}

	if err = cmd.Start(); err != nil {
		log.Fatal("Failed To Start The Plumbing")
	}

	go func() {
		defer close(ch)

		scanner := bufio.NewScanner(out)

		for scanner.Scan() {
			var j Journal
			if err := json.Unmarshal(scanner.Bytes(), &j); err != nil {
				log.Fatal("Failed To Reflect Journal: " + err.Error())
			}

			ch <- &j
		}
	}()

	return ch
}
