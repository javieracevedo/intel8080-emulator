package cpu

func (c *CPU) ADD_X(x Reg) {
	sum := c.REGISTERS[A] + c.REGISTERS[x]

	// Carry flag: check if result overflows 8 bits
	sum_16b := uint16(c.REGISTERS[A]) + uint16(c.REGISTERS[x])
	if sum_16b > 0xFF {
		c.SetFlag(CY)
	} else {
		c.ClearFlag(CY)
	}

	// Auxiliary carry: carry from bit 3 to bit 4
	if ((c.REGISTERS[A] & 0x0F) + (c.REGISTERS[x] & 0x0F)) > 0x0F {
		c.SetFlag(AC)
	} else {
		c.ClearFlag(AC)
	}

	// Zero flag: set if result is zero
	if sum == 0 {
		c.SetFlag(Z)
	} else {
		c.ClearFlag(Z)
	}

	// Sign flag: set if bit 7 of result is set
	if (sum & 0x80) != 0 {
		c.SetFlag(S)
	} else {
		c.ClearFlag(S)
	}

	// Parity flag: set if result has even parity
	if parity8(sum) {
		c.SetFlag(P)
	} else {
		c.ClearFlag(P)
	}

	c.REGISTERS[A] = sum
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
