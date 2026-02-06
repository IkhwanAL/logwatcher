# Log Watcher (Go Version)

## Description

Log Watcher, It Will Print The Log in The Log Text View Using Golang and tview

## Current State

- Able To View `journalctl` cmd
- Add TUI 

## Need To Add

- Mode Between Search Input And View And Scroll Input
- The Log TextView should be able to Scroll
- Able To Search The Log
- Add Checkbox To See which Log to be Shown in TextView
- Might As Well Add Pipe function where when type `any log cmd | log-watcher` so It fetch any log from it and show it into terminal

# Log Watcher (Go)

Log Watcher is a terminal-based log viewer written in Go using tview.
It focuses on making system logs readable, searchable, and interactive inside a TUI instead of dumping endless text into your terminal.

Think of it as a calm, structured window into noisy logs.

## Features (Current)

- View logs from journalctl
- Terminal User Interface (TUI) built with tview
- Logs rendered in a dedicated text view instead of raw stdout

## Planned Features

- Input modes
  - Switch between log viewing, scrolling, and search input
- Scrollable log view
  - Smooth scrolling through long log output
- Log search
  - Filter or highlight matching log entries
- Log source selection
  - Checkbox-based selection for which logs are shown
- Pipe support
  - Allow usage like:
  ```sh
  any-log-command | log-watcher
  ```
  Display piped input directly in the TUI

## Motivation

System logs are essential but unpleasant to work with in raw form.
This project explores how far a simple Go + TUI stack can go in turning logs into something navigable instead of overwhelming.

It’s also a hands-on experiment with:

- Go concurrency
- Terminal UI design
- Handling streaming data cleanly

Tech Stack

- Go
- tview
