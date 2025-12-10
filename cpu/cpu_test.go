package cpu

import (
	"8080/memory"
	"fmt"
	"testing"
)

// ANSI color codes for test output
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorBold   = "\033[1m"
)

// Helper to format a pass message
func logPass(t *testing.T, opcode byte, name string, details string) {
	t.Logf("%s%s✓ PASS%s [0x%02X] %-12s %s",
		colorBold, colorGreen, colorReset, opcode, name, details)
}

// Helper to format a fail message
func logFail(t *testing.T, opcode byte, name string, expected, got string) {
	t.Errorf("%s%s✗ FAIL%s [0x%02X] %-12s expected: %s, got: %s",
		colorBold, colorRed, colorReset, opcode, name, expected, got)
}

// Helper to get register name
func regName(r Reg) string {
	return REGISTERS_NAMES[r]
}

// Helper function to create and initialize a CPU for testing
func newTestCPU() *CPU {
	c := &CPU{}
	c.Init()
	return c
}

// Helper function to reset registers to known values
func resetRegisters() {
	// Reset all registers to 0
	for i := range REGISTERS {
		REGISTERS[i] = 0
	}
}

// Helper function to clear memory
func clearMemory() {
	for i := range memory.MEMORY {
		memory.MEMORY[i] = 0
	}
}

// =============================================================================
// CPU INITIALIZATION TESTS
// =============================================================================

func TestCPU_Init(t *testing.T) {
	c := &CPU{}
	c.Init()

	if c.CyclesCount != 0 {
		t.Errorf("CyclesCount should be 0 after Init, got %d", c.CyclesCount)
	}

	// Verify ALL cycle values in the table
	expectedCycles := map[byte]byte{
		// MOV B,x (0x40-0x47)
		0x40: 5, 0x41: 5, 0x42: 5, 0x43: 5, 0x44: 5, 0x45: 5, 0x46: 7, 0x47: 5,
		// MOV C,x (0x48-0x4F)
		0x48: 5, 0x49: 5, 0x4A: 5, 0x4B: 5, 0x4C: 5, 0x4D: 5, 0x4E: 7, 0x4F: 5,
		// MOV D,x (0x50-0x57)
		0x50: 5, 0x51: 5, 0x52: 5, 0x53: 5, 0x54: 5, 0x55: 5, 0x56: 7, 0x57: 5,
		// MOV E,x (0x58-0x5F)
		0x58: 5, 0x59: 5, 0x5A: 5, 0x5B: 5, 0x5C: 5, 0x5D: 5, 0x5E: 7, 0x5F: 5,
		// MOV H,x (0x60-0x67)
		0x60: 5, 0x61: 5, 0x62: 5, 0x63: 5, 0x64: 5, 0x65: 5, 0x66: 7, 0x67: 5,
		// MOV L,x (0x68-0x6F)
		0x68: 5, 0x69: 5, 0x6A: 5, 0x6B: 5, 0x6C: 5, 0x6D: 5, 0x6E: 7, 0x6F: 5,
		// MOV M,x and MOV A,x (0x70-0x7F)
		0x70: 7, 0x71: 7, 0x72: 7, 0x73: 7, 0x74: 7, 0x75: 7, 0x76: 7, 0x77: 7,
		0x78: 5, 0x79: 5, 0x7A: 5, 0x7B: 5, 0x7C: 5, 0x7D: 5, 0x7E: 7, 0x7F: 5,
	}

	for opcode, expected := range expectedCycles {
		if c.CYCLES_TABLE[opcode] != expected {
			t.Errorf("CYCLES_TABLE[0x%02X] = %d, want %d", opcode, c.CYCLES_TABLE[opcode], expected)
		}
	}
}

// =============================================================================
// NOP INSTRUCTION TESTS
// =============================================================================

func TestNOP(t *testing.T) {
	c := newTestCPU()
	resetRegisters()

	// Set some initial register values
	REGISTERS[A] = 0x55
	REGISTERS[B] = 0xAA

	// Execute NOP
	c.Execute(0x00)

	// NOP should not change any registers
	if REGISTERS[A] != 0x55 {
		t.Errorf("NOP modified register A: got 0x%02X, want 0x55", REGISTERS[A])
	}
	if REGISTERS[B] != 0xAA {
		t.Errorf("NOP modified register B: got 0x%02X, want 0xAA", REGISTERS[B])
	}

	// Verify NOP doesn't add any cycles (0x00 is not in the cycles table)
	if c.CyclesCount != 0 {
		t.Errorf("NOP should add 0 cycles, got %d", c.CyclesCount)
	}
}

// =============================================================================
// MOV REGISTER-TO-REGISTER TESTS
// =============================================================================

func TestMOV_RegisterToRegister(t *testing.T) {
	tests := []struct {
		name     string
		opcode   byte
		src      Reg
		dst      Reg
		srcValue byte
	}{
		// MOV B,x instructions
		{"MOV B,B", 0x40, B, B, 0x11},
		{"MOV B,C", 0x41, C, B, 0x22},
		{"MOV B,D", 0x42, D, B, 0x33},
		{"MOV B,E", 0x43, E, B, 0x44},
		{"MOV B,H", 0x44, H, B, 0x55},
		{"MOV B,L", 0x45, L, B, 0x66},
		{"MOV B,A", 0x47, A, B, 0x77},

		// MOV C,x instructions
		{"MOV C,B", 0x48, B, C, 0x11},
		{"MOV C,C", 0x49, C, C, 0x22},
		{"MOV C,D", 0x4A, D, C, 0x33},
		{"MOV C,E", 0x4B, E, C, 0x44},
		{"MOV C,H", 0x4C, H, C, 0x55},
		{"MOV C,L", 0x4D, L, C, 0x66},
		{"MOV C,A", 0x4F, A, C, 0x77},

		// MOV D,x instructions
		{"MOV D,B", 0x50, B, D, 0x11},
		{"MOV D,C", 0x51, C, D, 0x22},
		{"MOV D,D", 0x52, D, D, 0x33},
		{"MOV D,E", 0x53, E, D, 0x44},
		{"MOV D,H", 0x54, H, D, 0x55},
		{"MOV D,L", 0x55, L, D, 0x66},
		{"MOV D,A", 0x57, A, D, 0x77},

		// MOV E,x instructions
		{"MOV E,B", 0x58, B, E, 0x11},
		{"MOV E,C", 0x59, C, E, 0x22},
		{"MOV E,D", 0x5A, D, E, 0x33},
		{"MOV E,E", 0x5B, E, E, 0x44},
		{"MOV E,H", 0x5C, H, E, 0x55},
		{"MOV E,L", 0x5D, L, E, 0x66},
		{"MOV E,A", 0x5F, A, E, 0x77},

		// MOV H,x instructions
		{"MOV H,B", 0x60, B, H, 0x11},
		{"MOV H,C", 0x61, C, H, 0x22},
		{"MOV H,D", 0x62, D, H, 0x33},
		{"MOV H,E", 0x63, E, H, 0x44},
		{"MOV H,H", 0x64, H, H, 0x55},
		{"MOV H,L", 0x65, L, H, 0x66},
		{"MOV H,A", 0x67, A, H, 0x77},

		// MOV L,x instructions
		{"MOV L,B", 0x68, B, L, 0x11},
		{"MOV L,C", 0x69, C, L, 0x22},
		{"MOV L,D", 0x6A, D, L, 0x33},
		{"MOV L,E", 0x6B, E, L, 0x44},
		{"MOV L,H", 0x6C, H, L, 0x55},
		{"MOV L,L", 0x6D, L, L, 0x66},
		{"MOV L,A", 0x6F, A, L, 0x77},

		// MOV A,x instructions
		{"MOV A,B", 0x78, B, A, 0x11},
		{"MOV A,C", 0x79, C, A, 0x22},
		{"MOV A,D", 0x7A, D, A, 0x33},
		{"MOV A,E", 0x7B, E, A, 0x44},
		{"MOV A,H", 0x7C, H, A, 0x55},
		{"MOV A,L", 0x7D, L, A, 0x66},
		{"MOV A,A", 0x7F, A, A, 0x77},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			// Set up source register with test value
			REGISTERS[tt.src] = tt.srcValue

			initialCycles := c.CyclesCount
			c.Execute(tt.opcode)

			// Verify the destination register received the source value
			passed := true
			if REGISTERS[tt.dst] != tt.srcValue {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("%s=0x%02X", regName(tt.dst), tt.srcValue),
					fmt.Sprintf("%s=0x%02X", regName(tt.dst), REGISTERS[tt.dst]))
				passed = false
			}

			// Verify cycles were added (5 cycles for reg-to-reg MOV)
			expectedCycles := initialCycles + 5
			if c.CyclesCount != expectedCycles {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("cycles=%d", expectedCycles),
					fmt.Sprintf("cycles=%d", c.CyclesCount))
				passed = false
			}

			if passed {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("%s ← %s (0x%02X)", regName(tt.dst), regName(tt.src), tt.srcValue))
			}
		})
	}
}

// =============================================================================
// MOV FROM MEMORY (MOV_X_M) TESTS
// =============================================================================

func TestMOV_X_M(t *testing.T) {
	tests := []struct {
		name   string
		opcode byte
		dst    Reg
	}{
		{"MOV B,M", 0x46, B},
		{"MOV C,M", 0x4E, C},
		{"MOV D,M", 0x56, D},
		{"MOV E,M", 0x5E, E},
		{"MOV H,M", 0x66, H},
		{"MOV L,M", 0x6E, L},
		{"MOV A,M", 0x7E, A},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			// Set up HL to point to our test address
			REGISTERS[H] = 0x12
			REGISTERS[L] = 0x34

			// Put test value in memory at HL address
			memory.MEMORY[0x1234] = 0xAB

			initialCycles := c.CyclesCount
			c.Execute(tt.opcode)

			// Verify the destination register received the value from memory
			passed := true
			if REGISTERS[tt.dst] != 0xAB {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("%s=0xAB", regName(tt.dst)),
					fmt.Sprintf("%s=0x%02X", regName(tt.dst), REGISTERS[tt.dst]))
				passed = false
			}

			// Verify cycles were added (7 cycles for MOV r,M)
			expectedCycles := initialCycles + 7
			if c.CyclesCount != expectedCycles {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("cycles=%d", expectedCycles),
					fmt.Sprintf("cycles=%d", c.CyclesCount))
				passed = false
			}

			if passed {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("%s ← (HL=0x1234) = 0xAB", regName(tt.dst)))
			}
		})
	}
}

func TestMOV_X_M_AddressCalculation(t *testing.T) {
	c := newTestCPU()
	resetRegisters()
	clearMemory()

	// Set up HL to point to address 0xABCD
	REGISTERS[H] = 0xAB
	REGISTERS[L] = 0xCD

	// Put test value at that address
	testValue := byte(0xDE)
	memory.MEMORY[0xABCD] = testValue

	c.Execute(0x7E) // MOV A,M

	if REGISTERS[A] != testValue {
		t.Errorf("MOV A,M: address calculation failed, got 0x%02X, want 0x%02X",
			REGISTERS[A], testValue)
	}
}

// =============================================================================
// MOV TO MEMORY (MOV_M_X) TESTS
// =============================================================================

func TestMOV_M_X(t *testing.T) {
	tests := []struct {
		name     string
		opcode   byte
		src      Reg
		srcIndex int
	}{
		{"MOV M,B", 0x70, B, 0},
		{"MOV M,C", 0x71, C, 1},
		{"MOV M,D", 0x72, D, 2},
		{"MOV M,E", 0x73, E, 3},
		{"MOV M,H", 0x74, H, 4},
		{"MOV M,L", 0x75, L, 5},
		{"MOV M,A", 0x77, A, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			// Set up HL to point to our test address
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00

			// Set source register to a known test value
			testValue := byte(0x42 + tt.srcIndex)
			REGISTERS[tt.src] = testValue

			// Calculate the actual address after setting the source register
			// (important for MOV M,H and MOV M,L where src affects HL)
			expectedAddr := (uint16(REGISTERS[H]) << 8) | uint16(REGISTERS[L])

			initialCycles := c.CyclesCount
			c.Execute(tt.opcode)

			// Verify memory at the calculated HL address received the register value
			passed := true
			if memory.MEMORY[expectedAddr] != testValue {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("(0x%04X)=0x%02X", expectedAddr, testValue),
					fmt.Sprintf("(0x%04X)=0x%02X", expectedAddr, memory.MEMORY[expectedAddr]))
				passed = false
			}

			// Verify cycles were added (7 cycles for MOV M,r)
			expectedCycles := initialCycles + 7
			if c.CyclesCount != expectedCycles {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("cycles=%d", expectedCycles),
					fmt.Sprintf("cycles=%d", c.CyclesCount))
				passed = false
			}

			if passed {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("(HL=0x%04X) ← %s (0x%02X)", expectedAddr, regName(tt.src), testValue))
			}
		})
	}
}

func TestMOV_M_X_AddressCalculation(t *testing.T) {
	c := newTestCPU()
	resetRegisters()
	clearMemory()

	// Set up HL to point to address 0xBEEF
	REGISTERS[H] = 0xBE
	REGISTERS[L] = 0xEF
	REGISTERS[A] = 0x99

	c.Execute(0x77) // MOV M,A

	if memory.MEMORY[0xBEEF] != 0x99 {
		t.Errorf("MOV M,A: address calculation failed, got 0x%02X, want 0x99",
			memory.MEMORY[0xBEEF])
	}
}

// =============================================================================
// CYCLES COUNTING TESTS
// =============================================================================

func TestCyclesCounting(t *testing.T) {
	c := newTestCPU()

	if c.CyclesCount != 0 {
		t.Errorf("Initial CyclesCount = %d, want 0", c.CyclesCount)
	}

	// Execute a series of instructions and verify total cycles
	// MOV B,C (5 cycles) + MOV B,M (7 cycles) + MOV M,A (7 cycles) = 19 cycles
	c.Execute(0x41) // MOV B,C - 5 cycles
	c.Execute(0x46) // MOV B,M - 7 cycles
	c.Execute(0x77) // MOV M,A - 7 cycles

	if c.CyclesCount != 19 {
		t.Errorf("CyclesCount after 3 instructions = %d, want 19", c.CyclesCount)
	}
}

func TestCycles_RegToReg(t *testing.T) {
	// All register-to-register MOV instructions should take 5 cycles
	regToRegOpcodes := []byte{
		0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, // MOV B,x
		0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, 0x4F, // MOV C,x
		0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, // MOV D,x
		0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5F, // MOV E,x
		0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, // MOV H,x
		0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6F, // MOV L,x
		0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7F, // MOV A,x
	}

	for _, opcode := range regToRegOpcodes {
		c := newTestCPU()
		c.Execute(opcode)

		if c.CyclesCount != 5 {
			t.Errorf("Opcode 0x%02X: CyclesCount = %d, want 5", opcode, c.CyclesCount)
		}
	}
}

func TestCycles_MemoryOperations(t *testing.T) {
	// All memory-based MOV instructions should take 7 cycles
	memOpcodes := []byte{
		0x46, 0x4E, 0x56, 0x5E, 0x66, 0x6E, 0x7E, // MOV r,M
		0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77, // MOV M,r
	}

	for _, opcode := range memOpcodes {
		c := newTestCPU()
		clearMemory()
		c.Execute(opcode)

		if c.CyclesCount != 7 {
			t.Errorf("Opcode 0x%02X: CyclesCount = %d, want 7", opcode, c.CyclesCount)
		}
	}
}

// =============================================================================
// MOV FUNCTION UNIT TESTS
// =============================================================================

func TestMOV_DirectCall(t *testing.T) {
	resetRegisters()

	REGISTERS[A] = 0x42
	REGISTERS[B] = 0x00

	MOV(B, A)

	if REGISTERS[B] != 0x42 {
		t.Errorf("MOV(B, A): B = 0x%02X, want 0x42", REGISTERS[B])
	}

	// Source should be unchanged
	if REGISTERS[A] != 0x42 {
		t.Errorf("MOV(B, A): A should remain 0x42, got 0x%02X", REGISTERS[A])
	}
}

func TestMOV_SameRegister(t *testing.T) {
	resetRegisters()

	REGISTERS[A] = 0xFF
	MOV(A, A)

	if REGISTERS[A] != 0xFF {
		t.Errorf("MOV(A, A): A = 0x%02X, want 0xFF", REGISTERS[A])
	}
}

// =============================================================================
// MOV_X_M FUNCTION UNIT TESTS
// =============================================================================

func TestMOV_X_M_DirectCall(t *testing.T) {
	resetRegisters()
	clearMemory()

	// Set HL to point to address 0x1234
	REGISTERS[H] = 0x12
	REGISTERS[L] = 0x34

	// Put a value at that address
	memory.MEMORY[0x1234] = 0xBE

	MOV_X_M(A)

	if REGISTERS[A] != 0xBE {
		t.Errorf("MOV_X_M(A): A = 0x%02X, want 0xBE", REGISTERS[A])
	}
}

func TestMOV_X_M_ZeroAddress(t *testing.T) {
	resetRegisters()
	clearMemory()

	// Set HL to point to address 0x0000
	REGISTERS[H] = 0x00
	REGISTERS[L] = 0x00

	memory.MEMORY[0x0000] = 0xCD

	MOV_X_M(B)

	if REGISTERS[B] != 0xCD {
		t.Errorf("MOV_X_M(B) at address 0x0000: B = 0x%02X, want 0xCD", REGISTERS[B])
	}
}

func TestMOV_X_M_HighAddress(t *testing.T) {
	resetRegisters()
	clearMemory()

	// Set HL to point to a higher address 0x8000
	// Note: Memory size is 64*1014 = 64896 bytes (0xFD80), so using 0x8000 which is within bounds
	REGISTERS[H] = 0x80
	REGISTERS[L] = 0x00

	memory.MEMORY[0x8000] = 0x99

	MOV_X_M(C)

	if REGISTERS[C] != 0x99 {
		t.Errorf("MOV_X_M(C) at address 0x8000: C = 0x%02X, want 0x99", REGISTERS[C])
	}
}

// =============================================================================
// MOV_M_X FUNCTION UNIT TESTS
// =============================================================================

func TestMOV_M_X_DirectCall(t *testing.T) {
	resetRegisters()
	clearMemory()

	// Set HL to point to address 0x2000
	REGISTERS[H] = 0x20
	REGISTERS[L] = 0x00

	// Set register value to store
	REGISTERS[A] = 0xEF

	MOV_M_X(A)

	if memory.MEMORY[0x2000] != 0xEF {
		t.Errorf("MOV_M_X(A): MEMORY[0x2000] = 0x%02X, want 0xEF", memory.MEMORY[0x2000])
	}
}

func TestMOV_M_X_ZeroAddress(t *testing.T) {
	resetRegisters()
	clearMemory()

	REGISTERS[H] = 0x00
	REGISTERS[L] = 0x00
	REGISTERS[D] = 0x12

	MOV_M_X(D)

	if memory.MEMORY[0x0000] != 0x12 {
		t.Errorf("MOV_M_X(D) at address 0x0000: MEMORY[0x0000] = 0x%02X, want 0x12",
			memory.MEMORY[0x0000])
	}
}

func TestMOV_M_X_HighAddress(t *testing.T) {
	resetRegisters()
	clearMemory()

	// Note: Memory size is 64*1014 = 64896 bytes (0xFD80), so using 0xFD00 which is within bounds
	REGISTERS[H] = 0xFD
	REGISTERS[L] = 0x00
	REGISTERS[E] = 0x88

	MOV_M_X(E)

	if memory.MEMORY[0xFD00] != 0x88 {
		t.Errorf("MOV_M_X(E) at address 0xFD00: MEMORY[0xFD00] = 0x%02X, want 0x88",
			memory.MEMORY[0xFD00])
	}
}

// =============================================================================
// REGISTER CONSTANTS TESTS
// =============================================================================

func TestRegisterConstants(t *testing.T) {
	// Verify register constants have expected values
	if B != 0 {
		t.Errorf("B constant = %d, want 0", B)
	}
	if C != 1 {
		t.Errorf("C constant = %d, want 1", C)
	}
	if D != 2 {
		t.Errorf("D constant = %d, want 2", D)
	}
	if E != 3 {
		t.Errorf("E constant = %d, want 3", E)
	}
	if H != 4 {
		t.Errorf("H constant = %d, want 4", H)
	}
	if L != 5 {
		t.Errorf("L constant = %d, want 5", L)
	}
	if A != 6 {
		t.Errorf("A constant = %d, want 6", A)
	}
}

func TestRegisterNames(t *testing.T) {
	expectedNames := []string{"B", "C", "D", "E", "H", "L", "A"}

	for i, expected := range expectedNames {
		if REGISTERS_NAMES[i] != expected {
			t.Errorf("REGISTERS_NAMES[%d] = %s, want %s", i, REGISTERS_NAMES[i], expected)
		}
	}
}

// =============================================================================
// EDGE CASES AND BOUNDARY TESTS
// =============================================================================

func TestMOV_AllZeros(t *testing.T) {
	resetRegisters()
	clearMemory()

	// All registers are 0, move them around
	MOV(A, B)
	MOV(B, C)
	MOV(C, D)

	// All should still be 0
	for i, v := range REGISTERS {
		if v != 0 {
			t.Errorf("REGISTERS[%d] = 0x%02X after zero moves, want 0x00", i, v)
		}
	}
}

func TestMOV_AllFF(t *testing.T) {
	resetRegisters()

	// Set all registers to 0xFF
	for i := range REGISTERS {
		REGISTERS[i] = 0xFF
	}

	// Move between registers
	MOV(A, B)
	MOV(B, C)
	MOV(C, D)

	// All should still be 0xFF
	for i, v := range REGISTERS {
		if v != 0xFF {
			t.Errorf("REGISTERS[%d] = 0x%02X after 0xFF moves, want 0xFF", i, v)
		}
	}
}

func TestMOV_X_M_ReadZeroFromMemory(t *testing.T) {
	resetRegisters()
	clearMemory()

	REGISTERS[H] = 0x10
	REGISTERS[L] = 0x00
	REGISTERS[A] = 0xFF // Set A to non-zero

	// Memory is cleared, so reading should give 0
	MOV_X_M(A)

	if REGISTERS[A] != 0x00 {
		t.Errorf("MOV_X_M reading cleared memory: A = 0x%02X, want 0x00", REGISTERS[A])
	}
}

func TestMOV_M_X_WriteZeroToMemory(t *testing.T) {
	resetRegisters()
	clearMemory()

	REGISTERS[H] = 0x10
	REGISTERS[L] = 0x00
	memory.MEMORY[0x1000] = 0xFF // Set memory to non-zero

	REGISTERS[A] = 0x00
	MOV_M_X(A)

	if memory.MEMORY[0x1000] != 0x00 {
		t.Errorf("MOV_M_X writing zero: MEMORY[0x1000] = 0x%02X, want 0x00",
			memory.MEMORY[0x1000])
	}
}

// =============================================================================
// SEQUENTIAL OPERATION TESTS
// =============================================================================

func TestSequentialMoves(t *testing.T) {
	resetRegisters()

	// Create a chain: A -> B -> C -> D
	REGISTERS[A] = 0x42

	MOV(B, A) // B = 0x42
	MOV(C, B) // C = 0x42
	MOV(D, C) // D = 0x42

	if REGISTERS[B] != 0x42 {
		t.Errorf("Sequential move chain: B = 0x%02X, want 0x42", REGISTERS[B])
	}
	if REGISTERS[C] != 0x42 {
		t.Errorf("Sequential move chain: C = 0x%02X, want 0x42", REGISTERS[C])
	}
	if REGISTERS[D] != 0x42 {
		t.Errorf("Sequential move chain: D = 0x%02X, want 0x42", REGISTERS[D])
	}
}

func TestMemoryRoundTrip(t *testing.T) {
	resetRegisters()
	clearMemory()

	// Store a value to memory, then load it back to a different register
	REGISTERS[H] = 0x20
	REGISTERS[L] = 0x00
	REGISTERS[A] = 0x55

	MOV_M_X(A) // Store A (0x55) to memory at 0x2000

	// Verify it's in memory
	if memory.MEMORY[0x2000] != 0x55 {
		t.Errorf("Store phase: MEMORY[0x2000] = 0x%02X, want 0x55", memory.MEMORY[0x2000])
	}

	// Now load it back to B
	REGISTERS[B] = 0x00
	MOV_X_M(B)

	if REGISTERS[B] != 0x55 {
		t.Errorf("Load phase: B = 0x%02X, want 0x55", REGISTERS[B])
	}
}

// =============================================================================
// UNIMPLEMENTED OPCODE TESTS
// =============================================================================

func TestUnimplementedOpcode(t *testing.T) {
	c := newTestCPU()

	// Execute an opcode that's not implemented (e.g., 0x01 - LXI B)
	// This should just print a message but not panic
	// We can't easily test the output, but we can verify it doesn't crash
	c.Execute(0x01) // Should print "NOT IMPLEMENTED" but not crash
	c.Execute(0xFF) // Another unimplemented opcode
}
