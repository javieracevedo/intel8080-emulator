package main

import "8080/cpu"

var testInstructions = []byte{
	0x00,
	0x40,
	0x41,
	0x42,
	0x43,
	0x44,
	0x45,
	0x46,
	0x47,
}

func main() {
	c := cpu.CPU{}

	for _, v := range testInstructions {
		c.Execute(v)
	}

	c.DebugRegisters()
}


