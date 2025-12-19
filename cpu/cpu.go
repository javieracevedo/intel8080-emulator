package cpu

import (
	"fmt"
)

type CPU struct {
	Flags        byte
	PC           uint16
	CYCLES_TABLE [256]byte
	CyclesCount  uint
	REGISTERS    [7]byte
}

type Reg byte

const (
	B Reg = iota
	C
	D
	E
	H
	L
	A
)

func (c *CPU) Init(regs ...[7]byte) {
	c.CyclesCount = 0
	if len(regs) > 0 {
		c.REGISTERS = regs[0]
	} else {
		c.REGISTERS = [7]byte{}
	}
	c.CYCLES_TABLE = [256]byte{
		0x40: 5,
		0x41: 5,
		0x42: 5,
		0x43: 5,
		0x44: 5,
		0x45: 5,
		0x46: 7,
		0x47: 5,
		0x48: 5,
		0x49: 5,
		0x4A: 5,
		0x4B: 5,
		0x4C: 5,
		0x4D: 5,
		0x4E: 7,
		0x4F: 5,

		0x50: 5,
		0x51: 5,
		0x52: 5,
		0x53: 5,
		0x54: 5,
		0x55: 5,
		0x56: 7,
		0x57: 5,
		0x58: 5,
		0x59: 5,
		0x5A: 5,
		0x5B: 5,
		0x5C: 5,
		0x5D: 5,
		0x5E: 7,
		0x5F: 5,

		0x60: 5,
		0x61: 5,
		0x62: 5,
		0x63: 5,
		0x64: 5,
		0x65: 5,
		0x66: 7,
		0x67: 5,
		0x68: 5,
		0x69: 5,
		0x6A: 5,
		0x6B: 5,
		0x6C: 5,
		0x6D: 5,
		0x6E: 7,
		0x6F: 5,

		0x70: 7,
		0x71: 7,
		0x72: 7,
		0x73: 7,
		0x74: 7,
		0x75: 7,
		0x76: 7,
		0x77: 7,
		0x78: 5,
		0x79: 5,
		0x7A: 5,
		0x7B: 5,
		0x7C: 5,
		0x7D: 5,
		0x7E: 7,
		0x7F: 5,
	}
}

var REGISTERS_NAMES = [7]string{"B", "C", "D", "E", "H", "L", "A"}

func (c *CPU) DebugRegisters() {
	for i, v := range c.REGISTERS {
		fmt.Printf("%s: %02X\n", REGISTERS_NAMES[i], v)
	}
}
