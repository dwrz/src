package log

import (
	"log"
	"os"
)

const flags = log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC

type Logger struct {
	Debug *log.Logger
	Error *log.Logger
}

func New(f *os.File) *Logger {
	return &Logger{
		Debug: log.New(f, "DEBUG ", flags),
		Error: log.New(f, "ERROR ", flags),
	}
}
