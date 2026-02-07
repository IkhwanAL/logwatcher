package log

import (
	"io"
	"log"
	"os/exec"
)

func WatchJournal() io.ReadCloser {
	cmd := exec.Command("journalctl", "-f", "--no-tail", "--output=short")
	out, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalf("Failed To Pipe In To Log TUI ")
	}

	if err = cmd.Start(); err != nil {
		log.Fatalf("Failed To Start The Plumbing")
	}

	return out
}
