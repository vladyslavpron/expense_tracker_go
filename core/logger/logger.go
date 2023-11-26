package logger

import (
	"log"
)

func Log(message string) {
	log.Println("LOG: ", message)
}

func Warn(message string) {
	log.Println("WARNING: ", message)
}

func Err(message string) {
	log.Println("ERR: ", message)
}
