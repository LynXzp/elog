package elog

import (
	"log"
)

var minPrintableLevel int = 1

// SetLevel set minimal logging level
// higher level means more important message
// messages with smaller level will be ignored
// Fatal messages always will be printed and Panic after
func SetLevel(level int) {
	minPrintableLevel = level
}

func Msg(level int, msg string) {
	if minPrintableLevel > level {
		return
	}
	log.Println(msg)
}

func IfErr(level int, err error) {
	if minPrintableLevel > level {
		return
	}
	if err != nil {
		log.Println(err)
	}
}

func IfErrMsg(level int, err error, msg string) {
	if minPrintableLevel > level {
		return
	}
	if err != nil {
		log.Println(msg, ":", err)
	}
}

func IfErrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IfErrMsgFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, ":", err)
	}
}

func Fatal(msg string) {
	log.Fatal(msg)
}
