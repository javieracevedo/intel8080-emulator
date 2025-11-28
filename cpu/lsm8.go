package cpu
import "fmt"

func MOV(left Reg, right Reg) {
	REGISTERS[left] = REGISTERS[right]
}

