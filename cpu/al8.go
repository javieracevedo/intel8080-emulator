package cpu

import (
	"8080/memory"
)

func (c *CPU) ADD_X(x Reg) {
	sum := c.REGISTERS[A] + c.REGISTERS[x]

	c.SetAddConditionFlags(c.REGISTERS[A], c.REGISTERS[B])

	c.REGISTERS[A] = sum
}

func (c *CPU) ADD_M_X(x Reg) {
	MSB := c.REGISTERS[H]
	LSB := c.REGISTERS[L]
	var addr uint16 = (uint16(MSB) << 8) | uint16(LSB)

	sum := c.REGISTERS[x] + memory.MEMORY[addr]

	c.SetAddConditionFlags(c.REGISTERS[x], memory.MEMORY[addr])

	c.REGISTERS[A] = sum
}

func (c *CPU) ADC_X(x Reg) {
	var carry byte
	if c.IsSet(CY) {
		carry = 1
	}

	var ab_sum byte = c.REGISTERS[A] + c.REGISTERS[x]
	c.REGISTERS[A] = ab_sum + carry

	c.SetAddConditionFlags(ab_sum, carry)
}

func (c *CPU) SetAddConditionFlags(a byte, b byte) {
	// Carry flag: check if result overflows 8 bits
	sum_16b := uint16(a) + uint16(b)
	if sum_16b > 0xFF {
		c.SetFlag(CY)
	} else {
		c.ClearFlag(CY)
	}

	// Auxiliary carry: carry from bit 3 to bit 4
	if ((a & 0x0F) + (b & 0x0F)) > 0x0F {
		c.SetFlag(AC)
	} else {
		c.ClearFlag(AC)
	}

	// Zero flag: set if result is zero
	if byte(sum_16b) == 0 {
		c.SetFlag(Z)
	} else {
		c.ClearFlag(Z)
	}

	// Sign flag: set if bit 7 of result is set
	if (sum_16b & 0x80) != 0 {
		c.SetFlag(S)
	} else {
		c.ClearFlag(S)
	}

	// Parity flag: set if result has even parity
	if parity8(byte(sum_16b)) {
		c.SetFlag(P)
	} else {
		c.ClearFlag(P)
	}
}

func parity8(x byte) bool {
	count := 0
	for i := range 8 {
		if (x>>i)&1 == 1 {
			count++
		}
	}
	return count%2 == 0
}
