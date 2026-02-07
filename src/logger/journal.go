package logger

import (
	"bufio"
	"encoding/json"
	"log"
	"os/exec"
)

type Journal struct {
	MESSAGE              string
	PRIORITY             string
	__REALTIME_TIMESTAMP string
	_SYSTEMD_UNIT        string
	_COMM                string
	_TRANSPORT           string
	COREDUMP_SIGNAL_NAME string
	COREDUMP_COMM        string
	COREDUMP_PID         string
	_BOOT_ID             string
	rawLine              string
}

func (j *Journal) Parse() string {
	return j.MESSAGE
}

// Creating A Pipeline And Start Plumbing
func WatchJournal() <-chan Log {
	ch := make(chan Log)

	go func() {
		defer close(ch)
		cmd := exec.Command("journalctl", "-f", "--no-tail", "--output=json")
		out, err := cmd.StdoutPipe()

		if err != nil {
			log.Fatal("Failed To Pipe In To Log TUI ")
		}

		if err = cmd.Start(); err != nil {
			log.Fatal("Failed To Start The Plumbing")
		}

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
