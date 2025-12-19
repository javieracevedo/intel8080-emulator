package cpu

func (c *CPU) Execute(op byte) {
	c.CyclesCount += uint(c.CYCLES_TABLE[op])
	switch op {
	case 0x00:
		// NOP - no operation
	// MOVE INSTRUCTIONS
	case 0x40:
		c.MOV(B, B)
	case 0x41:
		c.MOV(B, C)
	case 0x42:
		c.MOV(B, D)
	case 0x43:
		c.MOV(B, E)
	case 0x44:
		c.MOV(B, H)
	case 0x45:
		c.MOV(B, L)
	case 0x46:
		c.MOV_X_M(B)
	case 0x47:
		c.MOV(B, A)
	case 0x48:
		c.MOV(C, B)
	case 0x49:
		c.MOV(C, C)
	case 0x4A:
		c.MOV(C, D)
	case 0x4B:
		c.MOV(C, E)
	case 0x4C:
		c.MOV(C, H)
	case 0x4D:
		c.MOV(C, L)
	case 0x4E:
		c.MOV_X_M(C)
	case 0x4F:
		c.MOV(C, A)
	case 0x50:
		c.MOV(D, B)
	case 0x51:
		c.MOV(D, C)
	case 0x52:
		c.MOV(D, D)
	case 0x53:
		c.MOV(D, E)
	case 0x54:
		c.MOV(D, H)
	case 0x55:
		c.MOV(D, L)
	case 0x56:
		c.MOV_X_M(D)
	case 0x57:
		c.MOV(D, A)
	case 0x58:
		c.MOV(E, B)
	case 0x59:
		c.MOV(E, C)
	case 0x5A:
		c.MOV(E, D)
	case 0x5B:
		c.MOV(E, E)
	case 0x5C:
		c.MOV(E, H)
	case 0x5D:
		c.MOV(E, L)
	case 0x5E:
		c.MOV_X_M(E)
	case 0x5F:
		c.MOV(E, A)
	case 0x60:
		c.MOV(H, B)
	case 0x61:
		c.MOV(H, C)
	case 0x62:
		c.MOV(H, D)
	case 0x63:
		c.MOV(H, E)
	case 0x64:
		c.MOV(H, H)
	case 0x65:
		c.MOV(H, L)
	case 0x66:
		c.MOV_X_M(H)
	case 0x67:
		c.MOV(H, A)
	case 0x68:
		c.MOV(L, B)
	case 0x69:
		c.MOV(L, C)
	case 0x6A:
		c.MOV(L, D)
	case 0x6B:
		c.MOV(L, E)
	case 0x6C:
		c.MOV(L, H)
	case 0x6D:
		c.MOV(L, L)
	case 0x6E:
		c.MOV_X_M(L)
	case 0x6F:
		c.MOV(L, A)
	case 0x70:
		c.MOV_M_X(B)
	case 0x71:
		c.MOV_M_X(C)
	case 0x72:
		c.MOV_M_X(D)
	case 0x73:
		c.MOV_M_X(E)
	case 0x74:
		c.MOV_M_X(H)
	case 0x75:
		c.MOV_M_X(L)
	case 0x77:
		c.MOV_M_X(A)
	case 0x78:
		c.MOV(A, B)
	case 0x79:
		c.MOV(A, C)
	case 0x7A:
		c.MOV(A, D)
	case 0x7B:
		c.MOV(A, E)
	case 0x7C:
		c.MOV(A, H)
	case 0x7D:
		c.MOV(A, L)
	case 0x7E:
		c.MOV_X_M(A)
	case 0x7F:
		c.MOV(A, A)
	// ADD INSTRUCTIONS
	case 0x80:
		c.ADD_X(B)
	case 0x81:
		c.ADD_X(C)
	}
}
