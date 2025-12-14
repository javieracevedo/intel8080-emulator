package cpu

func ADD_X(x Reg) {
	sum := REGISTERS[A] + REGISTERS[x]

	sum_16b := uint16(REGISTERS[A]) + uint16(REGISTERS[x])
	carry := (sum_16b%0x100)&0x100 != 0
	if carry {
		FLAGS[CY] = 1
	} else {
		FLAGS[CY] = 0
	}

	auxiliary_carry := ((REGISTERS[A] & 0x0F) + (REGISTERS[x] & 0x0F)) > 0x0F
	if auxiliary_carry {
		FLAGS[AC] = 1
	} else {
		FLAGS[AC] = 0
	}

	if sum == 0 {
		FLAGS[Z] = 1
	} else {
		FLAGS[Z] = 0
	}

	sign_bit := (sum & 0x80)
	if sign_bit == 1 {
		FLAGS[S] = 1
	} else {
		FLAGS[S] = 0
	}

	has_parity := parity8(sum)
	if has_parity {
		FLAGS[P] = 1
	} else {
		FLAGS[P] = 0
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
