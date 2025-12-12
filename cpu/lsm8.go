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

func ADD_X(x Reg) {
	result := REGISTERS[A] + REGISTERS[x]

	if result == 0 {
		FLAGS[F_Z] = 1
	} else {
		FLAGS[F_Z] = 0
	}

	msb := (result >> 7) & 1 // If the MSB is 1, then the result is a negative number
	FLAGS[F_S] = msb

	overflow := result < REGISTERS[x]
	if overflow {
		FLAGS[F_C] = 1
	} else {
		FLAGS[F_C] = 0
	}

	REGISTERS[A] = result
}
