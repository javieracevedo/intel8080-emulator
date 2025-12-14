package cpu

import (
	"8080/memory"
)

func MOV(left Reg, right Reg) {
	REGISTERS[left] = REGISTERS[right]
}

func MOV_X_M(x Reg) {
	MSB := REGISTERS[H]
	LSB := REGISTERS[L]
	addr := (uint16(MSB) << 8) | uint16(LSB)
	value := memory.MEMORY[addr]
	REGISTERS[x] = value
}

func MOV_M_X(x Reg) {
	MSB := REGISTERS[H]
	LSB := REGISTERS[L]
	addr := (uint16(MSB) << 8) | uint16(LSB)
	value := REGISTERS[x]
	memory.MEMORY[addr] = value
}
