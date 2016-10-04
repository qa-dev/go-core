package log

import (
	"github.com/fatih/color"
	"log"
)

func Info(format string, a ...interface{}) {
	color.Blue("INFO: ")
	log.Printf(format, a ...)
}

func Fatal(format string, a ...interface{}) {
	color.Red("FATAL: ")
	log.Fatalf(format, a ...)
}
