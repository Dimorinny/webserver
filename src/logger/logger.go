package logger

import (
	"fmt"
	"log"
	"time"
)

const (
	debugPrefix = "Debug: "
	errorPrefix = "Error: "
	LogFileName = "server.log"
)

var instance *log.Logger

func Init(logger *log.Logger) {
	instance = logger
}

func LogD(v ...interface{}) {
	if instance != nil {
		instance.Println(currentTime() + debugPrefix + fmt.Sprint(v...))
	}
}

func LogE(v ...interface{}) {
	if instance != nil {
		instance.Fatal(currentTime() + errorPrefix + fmt.Sprint(v...))
	}
}

func currentTime() string {
	return time.Now().Format("2006-01-02 15:04:05 ")
}