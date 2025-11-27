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

func MOV_C_B() {
	fmt.Println("DEBUG: called MOV_B_C")
	REGISTERS[C] = REGISTERS[B]
}

func MOV_C_C() {
	fmt.Println("DEBUG: called MOV_C_C")
	REGISTERS[C] = REGISTERS[C] // does nothing
}

func MOV_C_D() {
	fmt.Println("DEBUG: called MOV_C_D")
	REGISTERS[C] = REGISTERS[D]
}

func MOV_C_E() {
	fmt.Println("DEBUG: called MOV_C_E")
	REGISTERS[C] = REGISTERS[E]
}

func MOV_C_H() {
	fmt.Println("DEBUG: called MOV_C_E")
	REGISTERS[C] = REGISTERS[E]
}

func MOV_C_L() {
	fmt.Println("DEBUG: called MOV_C_L")
	REGISTERS[C] = REGISTERS[L]
}

func MOV_C_A() {
	fmt.Println("DEBUG: called MOV_C_A")
	REGISTERS[C] = REGISTERS[A]
}

func MOV_D_B() {
	fmt.Println("DEBUG: called MOV_D_B")
	REGISTERS[D] = REGISTERS[B]
}

func MOV_D_C() {
	fmt.Println("DEBUG: called MOV_D_B")
	REGISTERS[D] = REGISTERS[C]
}

func MOV_D_D() {
	fmt.Println("DEBUG: called MOV_D_B")
	REGISTERS[D] = REGISTERS[D]
}

func MOV_D_E() {
	fmt.Println("DEBUG: called MOV_D_B")
	REGISTERS[D] = REGISTERS[E]
}

func MOV_D_H() {
	fmt.Println("DEBUG: called MOV_D_B")
	REGISTERS[D] = REGISTERS[H]
}

func MOV_D_L() {
	fmt.Println("DEBUG: called MOV_D_B")
	REGISTERS[D] = REGISTERS[L]
}

func MOV_D_A() {
	fmt.Println("DEBUG: called MOV_D_A")
	REGISTERS[D] = REGISTERS[A]
}

/*func MOV_B_M() { }*/ // To be implemented when memory is implemented
