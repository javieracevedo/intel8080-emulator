package cpu
import "fmt"

func (c *CPU) Execute(op byte) {
	switch op {
	case 0x00:
		fmt.Printf("CPU->EXECUTE : OP CODE [%02X] NO OP\n", op)
	// MOVE INSTRUCTIONS
	case 0x40:
		MOV(B, B)
	case 0x41:
		MOV(B, C)
	case 0x42:
		MOV(B, D)
	case 0x43:
		MOV(B, E)
	case 0x44:
		MOV(B, H)
	case 0x45:
		MOV(B, L)
	case 0x46:
		MOV_X_M(B)
	case 0x47:
		MOV(B, A)
	case 0x48:
		MOV(C, B)
	case 0x49:
		MOV(C, C)
	case 0x4A:
		MOV(C, D)
	case 0x4B:
		MOV(C, E)
	case 0x4C:
		MOV(C, H)
	case 0x4D:
		MOV(C, L)
	case 0x4E:
		MOV_X_M(C)
	case 0x4F:
		MOV(C, A)
	case 0x50:
		MOV(D, B)
	case 0x51:
		MOV(D, C)
	case 0x52:
		MOV(D, D)
	case 0x53:
		MOV(D, E)
	case 0x54:
		MOV(D, H)
	case 0x55:
		MOV(D, L)
	case 0x56:
		MOV_X_M(D)
	case 0x57:
		MOV(D, A)
	case 0x58:
		MOV(E, B)
	case 0x59:
		MOV(E, C)
	case 0x5A:
		MOV(E, D)
	case 0x5B:
		MOV(E, E)
	case 0x5C:
		MOV(E, H)
	case 0x5D:
		MOV(E, H)
	case 0x5E:
		MOV_X_M(E)
	case 0x5F:
		MOV(E, A)
	case 0x60:
		MOV(H, B)
	case 0x61:
		MOV(H, C)
	case 0x62:
		MOV(H, D)
	case 0x63:
		MOV(H, E)
	case 0x64:
		MOV(H, H)
	case 0x65:
		MOV(H, L)
	case 0x66:
		MOV_X_M(H)
	case 0x67:
		MOV(H, A)
	case 0x68:
		MOV(L, B)
	case 0x69:
		MOV(L, C)
	case 0x6A:
		MOV(L, D)
	case 0x6B:
		MOV(L, E)
	case 0x6C:
		MOV(L, H)
	case 0x6D:
		MOV(L, L)
	case 0x6E:
		MOV_X_M(L)
	case 0x6F:
		MOV(L, A)
	case 0x70:
		MOV_M_X(B)
	case 0x71:
		MOV_M_X(C)
	case 0x72:
		MOV_M_X(D)
	case 0x73:
		MOV_M_X(E)
	case 0x74:
		MOV_M_X(H)
	case 0x75:
		MOV_M_X(L)
	case 0x77:
		MOV_M_X(A)
	case 0x78:
		MOV(A, B)
	case 0x79:
		MOV(A, C)
	case 0x7A:
		MOV(A, D)
	case 0x7B:
		MOV(A, E)
	case 0x7C:
		MOV(A, H)
	case 0x7D:
		MOV(A, L)
	case 0x7E:
		MOV_X_M(A)
	case 0x7F:
		MOV(A, A)
	default:
		fmt.Printf("CPU->EXECUTE : OP CODE [%02X] NOT IMPLEMENTED\n", op)
	}
}

func (c *CPU) DebugRegisters() {
	for i, v := range REGISTERS {
		fmt.Printf("%s: %02X\n", REGISTERS_NAMES[i], v)
	}
}
