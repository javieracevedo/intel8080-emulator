package cpu
import "fmt"

func (c *CPU) Execute(op byte) {
	switch op {
	case 0x00:
		fmt.Printf("CPU->EXECUTE : OP CODE [%02X] NO OP\n", op)
	// MOVE INSTRUCTIONS
	case 0x40:
		MOV_B_B()
	case 0x41:
		MOV_B_C()
	case 0x42:
		MOV_B_D()
	case 0x43:
		MOV_B_E()
	case 0x44:
		MOV_B_H()
	case 0x45:
		MOV_B_L()
	case 0x47:
		MOV_B_A()
	case 0x48:
		MOV_C_B(); 
	case 0x49:
		MOV_C_C();
	case 0x4A:
		MOV_C_D();
	case 0x4B:
		MOV_C_E();
	case 0x4C:
		MOV_C_H();
	case 0x4D:
		MOV_C_L();
	case 0x4F:
		MOV_C_A();
	default:
		fmt.Printf("CPU->EXECUTE : OP CODE [%02X] NOT IMPLEMENTED\n", op)
	}
}

func (c *CPU) DebugRegisters() {
	for i, v := range REGISTERS {
		fmt.Printf("%s: %02X\n", REGISTERS_NAMES[i], v)
	}
}
