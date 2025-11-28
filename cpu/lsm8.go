package cpu
import "fmt"
import "8080/memory"

func MOV(left Reg, right Reg) {
	fmt.Printf("MOV: %02X <- %02X\n", left, right)
	REGISTERS[left] = REGISTERS[right]
}

func MOV_X_M(x Reg) {
	MSB := REGISTERS[H]
	LSB := REGISTERS[L]
	addr := (uint16(MSB) << 8) | uint16(LSB)

	value := memory.MEMORY[addr]

	fmt.Printf("MOV_X_M: %02X <- %02X (addr: 0x%04X)\n", x, value, addr)

	REGISTERS[x] = value
}

