package cpu

import (
	"8080/memory"
)

func (c *CPU) MOV(left Reg, right Reg) {
	c.REGISTERS[left] = c.REGISTERS[right]
}

func (c *CPU) MOV_X_M(x Reg) {
	MSB := c.REGISTERS[H]
	LSB := c.REGISTERS[L]
	addr := (uint16(MSB) << 8) | uint16(LSB)
	value := memory.MEMORY[addr]
	c.REGISTERS[x] = value
}

func (c *CPU) MOV_M_X(x Reg) {
	MSB := c.REGISTERS[H]
	LSB := c.REGISTERS[L]
	addr := (uint16(MSB) << 8) | uint16(LSB)
	value := c.REGISTERS[x]
	memory.MEMORY[addr] = value
}
