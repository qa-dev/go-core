package log

import (
	"github.com/qa-dev/go-core/color"
	"log"
)

func Info(format string, a ...interface{}) {
	log.Printf("%s"+format, color.Color("INFO: ", color.Blue), a...)
}

func Fatal(format string, a ...interface{}) {
	log.Fatalf("%s "+format, color.Color("FATAL:", color.Red), a...)
}
