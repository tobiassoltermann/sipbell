// +build windows
package main

import (
	"log"
)

func Init(portNumber int) {
	log.Println("RELAY: Init ", portNumber, "")
}

func On() {
	log.Println("RELAY: On")
}

func Off() {
	log.Println("RELAY: Off")
}

func Close() {
	log.Println("RELAY: Close")
}
