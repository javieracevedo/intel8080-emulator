package cpu
import "fmt"

func MOV(left Reg, right Reg) {
	fmt.Printf("%02X <- %02X\n", left, right)
	REGISTERS[left] = REGISTERS[right]
}

