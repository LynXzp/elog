package elog

import "fmt"

var minPrintableLevel int = 1

// Set minimal logging level
// higher level means more important message
// messages with smaller level will be ignored
func SetLevel(level int){
	minPrintableLevel = level
}

func IfErr(level int, err error){
	if minPrintableLevel > level {
		return
	}
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
}

func IfErrMsg(level int, err error, msg string){
	if minPrintableLevel > level {
		return
	}
	if err != nil {
		fmt.Printf("%v: %v\n", msg, err.Error())
	}
}

func IfErrPanic(err error){
	if err != nil {
		str := fmt.Sprintf("%v\n", err.Error())
		fmt.Print(str)
		panic(str)
	}
}

func IfErrMsgPanic(err error, msg string){
	if err != nil {
		str := fmt.Sprintf("%v: %v\n", msg, err.Error())
		fmt.Print(str)
		panic(str)
	}
}
