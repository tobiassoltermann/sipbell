package main

import (
	"log"
	"os"
	"reflect"

	"github.com/stianeikeland/go-rpio"
)

var pin rpio.Pin

func Init(portNumber int) {
	err := rpio.Open()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	pin = rpio.Pin(portNumber)
	log.Println(reflect.TypeOf(pin))
	pin.Output()
	pin.High()
}

func On() {
	pin.Low()
}

func Off() {
	pin.High()
}

func Close() {
}
