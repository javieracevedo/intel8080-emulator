package cpu

func ADD_X(x Reg) {
	sum := REGISTERS[A] + REGISTERS[x]

	// Carry flag: check if result overflows 8 bits
	sum_16b := uint16(REGISTERS[A]) + uint16(REGISTERS[x])
	if sum_16b > 0xFF {
		SetFlag(CY)
	} else {
		ClearFlag(CY)
	}

	// Auxiliary carry: carry from bit 3 to bit 4
	if ((REGISTERS[A] & 0x0F) + (REGISTERS[x] & 0x0F)) > 0x0F {
		SetFlag(AC)
	} else {
		ClearFlag(AC)
	}

	// Zero flag: set if result is zero
	if sum == 0 {
		SetFlag(Z)
	} else {
		ClearFlag(Z)
	}

	// Sign flag: set if bit 7 of result is set
	if (sum & 0x80) != 0 {
		SetFlag(S)
	} else {
		ClearFlag(S)
	}

	// Parity flag: set if result has even parity
	if parity8(sum) {
		SetFlag(P)
	} else {
		ClearFlag(P)
	}

	REGISTERS[A] = sum
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
