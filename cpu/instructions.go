package cpu
import "fmt"

func MOV_B_B() {
	fmt.Println("DEBUG: called MOV_B_B")
	REGISTERS[B] = REGISTERS[B] // this does nothing really
}

func MOV_B_C() {
	fmt.Println("DEBUG: called MOV_B_C")
	REGISTERS[B] = REGISTERS[C]
}

func MOV_B_D() {
	fmt.Println("DEBUG: called MOV_B_D")
	REGISTERS[B] = REGISTERS[D]
}

func MOV_B_E() {
	fmt.Println("DEBUG: called MOV_B_E")
	REGISTERS[B] = REGISTERS[E]
}

func MOV_B_H() {
	fmt.Println("DEBUG: called MOV_B_H")
	REGISTERS[B] = REGISTERS[H]
}

func MOV_B_L() {
	fmt.Println("DEBUG: called MOV_B_L")
	REGISTERS[B] = REGISTERS[L]
}

func MOV_B_A() {
	fmt.Println("DEBUG: called MOV_B_A")
	REGISTERS[B] = REGISTERS[A]
}

/*func MOV_B_M() { }*/ // To be implemented when memory is implemented

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
	default:
		fmt.Printf("CPU->EXECUTE : OP CODE [%02X] NOT IMPLEMENTED\n", op)
	}
}

func (c *CPU) DebugRegisters() {
	for i, v := range REGISTERS {
		fmt.Printf("%s: %02X\n", REGISTERS_NAMES[i], v)
	}
}
