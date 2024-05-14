package logger

import (
	"io"
	"log"
	"os"
)

// NewSimpleLogger creates a new log.Logger that writes to a file.
// The file is created if it does not exist, and append to it if it does.
// The file is created with mode 0644.
// The logger also writes to os.Stdout.
// The logger prefixes each log entry with the current date and time and the file name and line number of the calling code.
// The flags argument defines the logging properties.
// The flags are Ldate, Ltime, and Lshortfile.
// The Ldate flag causes the logger to write the current date in the local time zone: 2009/01/23.
// The Ltime flag causes the logger to write the current time in the local time zone: 01:23:23.
// The Lshortfile flag causes the logger to write the file name and line number: logger.go:24.
func NewSimpleLogger(logFile string) *log.Logger {

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)

	}
	return log.New(io.MultiWriter(file, os.Stdout), "", log.Ldate|log.Ltime|log.Lshortfile)
}
