package main

import (
	"log"
)

func logInfo(s string) {
	log.SetPrefix("[INFO] ")
	log.Print(s)
}

func logError(s string) {
	log.SetPrefix("[ERROR] ")
	log.Print(s)
}

func logFatal(s string) {
	log.SetPrefix("[PANIC] ")
	log.Fatal(s)
}
