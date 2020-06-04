package main

import (
	"log"
	. "log"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

/*
BasicIO struct defines the IO pin struct, it contains
Pin a valid integer of 32-bit
Mode a valid integer
0 := input
1 :- output with low
2 := output with high
3 :- output with toggle
*/
type BasicIO struct {
	Pin  int
	Mode int
}

// SetPin is used to test mode of pin
func (io *BasicIO) SetPin() {

	pin := rpio.Pin(io.Pin)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	switch io.Mode {
	case 0:
		pin.Input()
		Println(pin.Read())
	case 1:
		pin.Output()
		pin.Low()

	case 2:
		pin.Output()
		pin.High()

	case 3:
		pin.Output()
		for x := 0; x < 20; x++ {
			pin.Toggle()
			time.Sleep(time.Second / 5)
		}
	default:
		log.Println("Inside default")
	}

}
