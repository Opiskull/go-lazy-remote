package main

import (
	"io"
	"log"
	"os"
)

func initLogging(fileName string) func() {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.Print("Logging to " + fileName)
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)
	return func() {
		if f != nil {
			f.Close()
		}
	}
}
