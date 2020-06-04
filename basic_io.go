package main

import (
	"log"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

type BasicIO struct {
	Pin      int
	IsOutput bool
	IsToggle bool
}

func (io *BasicIO) SetPin() {

	pin := rpio.Pin(io.Pin)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	if io.IsOutput {
		pin.Output()
	} else {
		pin.Input()
	}

	for x := 0; x < 20; x++ {
		if io.IsToggle {
			pin.Toggle()
		}
		time.Sleep(time.Second / 5)
	}

}
