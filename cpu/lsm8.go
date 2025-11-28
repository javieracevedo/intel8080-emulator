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

func MOV_M_X(x Reg) {
	MSB := REGISTERS[H]
	LSB := REGISTERS[L]
	addr := (uint16(MSB) << 8) | uint16(LSB)

	value := REGISTERS[x]

	memory.MEMORY[addr] = value

	fmt.Printf("MOV_M_X: %04X (addr: 0x%04X) <- %04X (register: %02X)\n", memory.MEMORY[addr], addr, value, x)
}

