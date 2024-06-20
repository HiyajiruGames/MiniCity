package log

import (
	"fmt"
	"log"
)

func Debug(message string) {
	print("DEBUG", message, nil)
}

func Info(message string) {
	print("INFO", message, nil)
}

func Warn(message string, err error) {
	print("WARN", message, err)
}

func Error(message string, err error) {
	print("ERROR", message, err)
}

func print(level string, message string, err error) {
	var str string
	if err != nil {
		str = fmt.Sprintf("[%s] %s >>> %s", level, message, err)
	} else {
		str = fmt.Sprintf("[%s] %s", level, message)
	}
	log.Print(str)
}
