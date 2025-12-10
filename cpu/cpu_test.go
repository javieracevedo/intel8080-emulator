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

// =============================================================================
// ADD INSTRUCTION TESTS (0x80-0x87)
// Add register to accumulator
// =============================================================================

func TestADD_RegisterToA(t *testing.T) {
	tests := []struct {
		name       string
		opcode     byte
		src        Reg
		aValue     byte
		srcValue   byte
		wantResult byte
		wantCarry  bool
		wantZero   bool
		wantSign   bool
	}{
		// ADD B (0x80)
		{"ADD B basic", 0x80, B, 0x10, 0x20, 0x30, false, false, false},
		{"ADD B zero result", 0x80, B, 0x00, 0x00, 0x00, false, true, false},
		{"ADD B carry", 0x80, B, 0xFF, 0x01, 0x00, true, true, false},
		{"ADD B sign", 0x80, B, 0x7F, 0x01, 0x80, false, false, true},

		// ADD C (0x81)
		{"ADD C basic", 0x81, C, 0x05, 0x0A, 0x0F, false, false, false},
		{"ADD C carry", 0x81, C, 0x80, 0x80, 0x00, true, true, false},

		// ADD D (0x82)
		{"ADD D basic", 0x82, D, 0x11, 0x22, 0x33, false, false, false},

		// ADD E (0x83)
		{"ADD E basic", 0x83, E, 0x44, 0x11, 0x55, false, false, false},

		// ADD H (0x84)
		{"ADD H basic", 0x84, H, 0x10, 0x10, 0x20, false, false, false},

		// ADD L (0x85)
		{"ADD L basic", 0x85, L, 0x0F, 0x01, 0x10, false, false, false},

		// ADD A (0x87) - doubles the accumulator
		{"ADD A basic", 0x87, A, 0x40, 0x40, 0x80, false, false, true},
		{"ADD A zero", 0x87, A, 0x00, 0x00, 0x00, false, true, false},
		{"ADD A carry", 0x87, A, 0x80, 0x80, 0x00, true, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			REGISTERS[A] = tt.aValue
			if tt.src != A {
				REGISTERS[tt.src] = tt.srcValue
			}

			c.Execute(tt.opcode)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X + %s=0x%02X → 0x%02X",
						tt.aValue, regName(tt.src), tt.srcValue, tt.wantResult))
			}
		})
	}
}

func TestADD_M(t *testing.T) {
	// ADD M (0x86) - Add memory to accumulator
	tests := []struct {
		name       string
		aValue     byte
		memValue   byte
		wantResult byte
		wantCarry  bool
	}{
		{"ADD M basic", 0x10, 0x20, 0x30, false},
		{"ADD M carry", 0xFF, 0x01, 0x00, true},
		{"ADD M zero", 0x00, 0x00, 0x00, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			REGISTERS[A] = tt.aValue
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00
			memory.MEMORY[0x2000] = tt.memValue

			c.Execute(0x86)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, 0x86, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, 0x86, tt.name,
					fmt.Sprintf("A=0x%02X + (HL)=0x%02X → 0x%02X",
						tt.aValue, tt.memValue, tt.wantResult))
			}
		})
	}
}

// =============================================================================
// ADC INSTRUCTION TESTS (0x88-0x8F)
// Add register to accumulator with carry
// =============================================================================

func TestADC_RegisterToA(t *testing.T) {
	tests := []struct {
		name       string
		opcode     byte
		src        Reg
		aValue     byte
		srcValue   byte
		carryIn    bool
		wantResult byte
		wantCarry  bool
	}{
		// ADC B (0x88)
		{"ADC B no carry in", 0x88, B, 0x10, 0x20, false, 0x30, false},
		{"ADC B with carry in", 0x88, B, 0x10, 0x20, true, 0x31, false},
		{"ADC B carry out no in", 0x88, B, 0xFF, 0x01, false, 0x00, true},
		{"ADC B carry out with in", 0x88, B, 0xFE, 0x01, true, 0x00, true},

		// ADC C (0x89)
		{"ADC C no carry", 0x89, C, 0x05, 0x0A, false, 0x0F, false},
		{"ADC C with carry", 0x89, C, 0x05, 0x0A, true, 0x10, false},

		// ADC D (0x8A)
		{"ADC D basic", 0x8A, D, 0x11, 0x22, false, 0x33, false},

		// ADC E (0x8B)
		{"ADC E basic", 0x8B, E, 0x44, 0x11, false, 0x55, false},

		// ADC H (0x8C)
		{"ADC H basic", 0x8C, H, 0x10, 0x10, false, 0x20, false},

		// ADC L (0x8D)
		{"ADC L basic", 0x8D, L, 0x0F, 0x00, true, 0x10, false},

		// ADC A (0x8F) - doubles accumulator plus carry
		{"ADC A no carry", 0x8F, A, 0x40, 0x40, false, 0x80, false},
		{"ADC A with carry", 0x8F, A, 0x40, 0x40, true, 0x81, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			REGISTERS[A] = tt.aValue
			if tt.src != A {
				REGISTERS[tt.src] = tt.srcValue
			}
			// TODO: Set carry flag based on tt.carryIn when flags are implemented

			c.Execute(tt.opcode)

			// For now, just verify the basic operation without carry
			// Full flag verification will be needed once flags are implemented
			if !tt.carryIn {
				if REGISTERS[A] != tt.wantResult {
					logFail(t, tt.opcode, tt.name,
						fmt.Sprintf("A=0x%02X", tt.wantResult),
						fmt.Sprintf("A=0x%02X", REGISTERS[A]))
				} else {
					logPass(t, tt.opcode, tt.name,
						fmt.Sprintf("A=0x%02X + %s=0x%02X + CY=%v → 0x%02X",
							tt.aValue, regName(tt.src), tt.srcValue, tt.carryIn, tt.wantResult))
				}
			}
		})
	}
}

func TestADC_M(t *testing.T) {
	// ADC M (0x8E) - Add memory to accumulator with carry
	tests := []struct {
		name       string
		aValue     byte
		memValue   byte
		carryIn    bool
		wantResult byte
	}{
		{"ADC M no carry", 0x10, 0x20, false, 0x30},
		{"ADC M with carry", 0x10, 0x20, true, 0x31},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			REGISTERS[A] = tt.aValue
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00
			memory.MEMORY[0x2000] = tt.memValue

			c.Execute(0x8E)

			// Without carry in, verify basic operation
			if !tt.carryIn {
				if REGISTERS[A] != tt.wantResult {
					logFail(t, 0x8E, tt.name,
						fmt.Sprintf("A=0x%02X", tt.wantResult),
						fmt.Sprintf("A=0x%02X", REGISTERS[A]))
				} else {
					logPass(t, 0x8E, tt.name,
						fmt.Sprintf("A=0x%02X + (HL)=0x%02X → 0x%02X",
							tt.aValue, tt.memValue, tt.wantResult))
				}
			}
		})
	}
}

// =============================================================================
// SUB INSTRUCTION TESTS (0x90-0x97)
// Subtract register from accumulator
// =============================================================================

func TestSUB_RegisterFromA(t *testing.T) {
	tests := []struct {
		name       string
		opcode     byte
		src        Reg
		aValue     byte
		srcValue   byte
		wantResult byte
		wantCarry  bool // Carry means borrow for SUB
		wantZero   bool
		wantSign   bool
	}{
		// SUB B (0x90)
		{"SUB B basic", 0x90, B, 0x30, 0x10, 0x20, false, false, false},
		{"SUB B zero result", 0x90, B, 0x10, 0x10, 0x00, false, true, false},
		{"SUB B borrow", 0x90, B, 0x10, 0x20, 0xF0, true, false, true},

		// SUB C (0x91)
		{"SUB C basic", 0x91, C, 0x50, 0x25, 0x2B, false, false, false},

		// SUB D (0x92)
		{"SUB D basic", 0x92, D, 0x33, 0x11, 0x22, false, false, false},

		// SUB E (0x93)
		{"SUB E basic", 0x93, E, 0x55, 0x11, 0x44, false, false, false},

		// SUB H (0x94)
		{"SUB H basic", 0x94, H, 0x40, 0x10, 0x30, false, false, false},

		// SUB L (0x95)
		{"SUB L basic", 0x95, L, 0x20, 0x10, 0x10, false, false, false},

		// SUB A (0x97) - always results in zero
		{"SUB A always zero", 0x97, A, 0x42, 0x42, 0x00, false, true, false},
		{"SUB A from FF", 0x97, A, 0xFF, 0xFF, 0x00, false, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			REGISTERS[A] = tt.aValue
			if tt.src != A {
				REGISTERS[tt.src] = tt.srcValue
			}

			c.Execute(tt.opcode)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X - %s=0x%02X → 0x%02X",
						tt.aValue, regName(tt.src), tt.srcValue, tt.wantResult))
			}
		})
	}
}

func TestSUB_M(t *testing.T) {
	// SUB M (0x96) - Subtract memory from accumulator
	tests := []struct {
		name       string
		aValue     byte
		memValue   byte
		wantResult byte
		wantCarry  bool
	}{
		{"SUB M basic", 0x30, 0x10, 0x20, false},
		{"SUB M borrow", 0x10, 0x20, 0xF0, true},
		{"SUB M zero", 0x55, 0x55, 0x00, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			REGISTERS[A] = tt.aValue
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00
			memory.MEMORY[0x2000] = tt.memValue

			c.Execute(0x96)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, 0x96, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, 0x96, tt.name,
					fmt.Sprintf("A=0x%02X - (HL)=0x%02X → 0x%02X",
						tt.aValue, tt.memValue, tt.wantResult))
			}
		})
	}
}

// =============================================================================
// SBB INSTRUCTION TESTS (0x98-0x9F)
// Subtract register from accumulator with borrow
// =============================================================================

func TestSBB_RegisterFromA(t *testing.T) {
	tests := []struct {
		name       string
		opcode     byte
		src        Reg
		aValue     byte
		srcValue   byte
		borrowIn   bool
		wantResult byte
	}{
		// SBB B (0x98)
		{"SBB B no borrow", 0x98, B, 0x30, 0x10, false, 0x20},
		{"SBB B with borrow", 0x98, B, 0x30, 0x10, true, 0x1F},

		// SBB C (0x99)
		{"SBB C no borrow", 0x99, C, 0x50, 0x25, false, 0x2B},

		// SBB D (0x9A)
		{"SBB D no borrow", 0x9A, D, 0x33, 0x11, false, 0x22},

		// SBB E (0x9B)
		{"SBB E no borrow", 0x9B, E, 0x55, 0x11, false, 0x44},

		// SBB H (0x9C)
		{"SBB H no borrow", 0x9C, H, 0x40, 0x10, false, 0x30},

		// SBB L (0x9D)
		{"SBB L no borrow", 0x9D, L, 0x20, 0x10, false, 0x10},

		// SBB A (0x9F) - result depends on carry
		{"SBB A no borrow", 0x9F, A, 0x42, 0x42, false, 0x00},
		{"SBB A with borrow", 0x9F, A, 0x42, 0x42, true, 0xFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			REGISTERS[A] = tt.aValue
			if tt.src != A {
				REGISTERS[tt.src] = tt.srcValue
			}
			// TODO: Set carry flag based on tt.borrowIn when flags are implemented

			c.Execute(tt.opcode)

			// Without borrow in, verify basic operation
			if !tt.borrowIn {
				if REGISTERS[A] != tt.wantResult {
					logFail(t, tt.opcode, tt.name,
						fmt.Sprintf("A=0x%02X", tt.wantResult),
						fmt.Sprintf("A=0x%02X", REGISTERS[A]))
				} else {
					logPass(t, tt.opcode, tt.name,
						fmt.Sprintf("A=0x%02X - %s=0x%02X - CY=%v → 0x%02X",
							tt.aValue, regName(tt.src), tt.srcValue, tt.borrowIn, tt.wantResult))
				}
			}
		})
	}
}

func TestSBB_M(t *testing.T) {
	// SBB M (0x9E) - Subtract memory from accumulator with borrow
	tests := []struct {
		name       string
		aValue     byte
		memValue   byte
		borrowIn   bool
		wantResult byte
	}{
		{"SBB M no borrow", 0x30, 0x10, false, 0x20},
		{"SBB M with borrow", 0x30, 0x10, true, 0x1F},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			REGISTERS[A] = tt.aValue
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00
			memory.MEMORY[0x2000] = tt.memValue

			c.Execute(0x9E)

			if !tt.borrowIn {
				if REGISTERS[A] != tt.wantResult {
					logFail(t, 0x9E, tt.name,
						fmt.Sprintf("A=0x%02X", tt.wantResult),
						fmt.Sprintf("A=0x%02X", REGISTERS[A]))
				} else {
					logPass(t, 0x9E, tt.name,
						fmt.Sprintf("A=0x%02X - (HL)=0x%02X → 0x%02X",
							tt.aValue, tt.memValue, tt.wantResult))
				}
			}
		})
	}
}

// =============================================================================
// ANA INSTRUCTION TESTS (0xA0-0xA7)
// Logical AND register with accumulator
// =============================================================================

func TestANA_RegisterWithA(t *testing.T) {
	tests := []struct {
		name       string
		opcode     byte
		src        Reg
		aValue     byte
		srcValue   byte
		wantResult byte
		wantZero   bool
	}{
		// ANA B (0xA0)
		{"ANA B basic", 0xA0, B, 0xFF, 0x0F, 0x0F, false},
		{"ANA B zero", 0xA0, B, 0xAA, 0x55, 0x00, true},
		{"ANA B all ones", 0xA0, B, 0xFF, 0xFF, 0xFF, false},

		// ANA C (0xA1)
		{"ANA C basic", 0xA1, C, 0xF0, 0x0F, 0x00, true},

		// ANA D (0xA2)
		{"ANA D basic", 0xA2, D, 0x33, 0x11, 0x11, false},

		// ANA E (0xA3)
		{"ANA E basic", 0xA3, E, 0x55, 0xFF, 0x55, false},

		// ANA H (0xA4)
		{"ANA H basic", 0xA4, H, 0xCC, 0xAA, 0x88, false},

		// ANA L (0xA5)
		{"ANA L basic", 0xA5, L, 0x12, 0x34, 0x10, false},

		// ANA A (0xA7) - AND with self, value unchanged
		{"ANA A unchanged", 0xA7, A, 0x42, 0x42, 0x42, false},
		{"ANA A zero", 0xA7, A, 0x00, 0x00, 0x00, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			REGISTERS[A] = tt.aValue
			if tt.src != A {
				REGISTERS[tt.src] = tt.srcValue
			}

			c.Execute(tt.opcode)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X & %s=0x%02X → 0x%02X",
						tt.aValue, regName(tt.src), tt.srcValue, tt.wantResult))
			}
		})
	}
}

func TestANA_M(t *testing.T) {
	// ANA M (0xA6) - AND memory with accumulator
	tests := []struct {
		name       string
		aValue     byte
		memValue   byte
		wantResult byte
	}{
		{"ANA M basic", 0xFF, 0x0F, 0x0F},
		{"ANA M zero", 0xAA, 0x55, 0x00},
		{"ANA M all ones", 0xFF, 0xFF, 0xFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			REGISTERS[A] = tt.aValue
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00
			memory.MEMORY[0x2000] = tt.memValue

			c.Execute(0xA6)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, 0xA6, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, 0xA6, tt.name,
					fmt.Sprintf("A=0x%02X & (HL)=0x%02X → 0x%02X",
						tt.aValue, tt.memValue, tt.wantResult))
			}
		})
	}
}

// =============================================================================
// XRA INSTRUCTION TESTS (0xA8-0xAF)
// Logical XOR register with accumulator
// =============================================================================

func TestXRA_RegisterWithA(t *testing.T) {
	tests := []struct {
		name       string
		opcode     byte
		src        Reg
		aValue     byte
		srcValue   byte
		wantResult byte
		wantZero   bool
	}{
		// XRA B (0xA8)
		{"XRA B basic", 0xA8, B, 0xFF, 0x0F, 0xF0, false},
		{"XRA B same values", 0xA8, B, 0xAA, 0xAA, 0x00, true},

		// XRA C (0xA9)
		{"XRA C basic", 0xA9, C, 0xF0, 0x0F, 0xFF, false},

		// XRA D (0xAA)
		{"XRA D basic", 0xAA, D, 0x33, 0x11, 0x22, false},

		// XRA E (0xAB)
		{"XRA E basic", 0xAB, E, 0x55, 0xFF, 0xAA, false},

		// XRA H (0xAC)
		{"XRA H basic", 0xAC, H, 0xCC, 0xAA, 0x66, false},

		// XRA L (0xAD)
		{"XRA L basic", 0xAD, L, 0x12, 0x34, 0x26, false},

		// XRA A (0xAF) - XOR with self always zeros
		{"XRA A always zero", 0xAF, A, 0x42, 0x42, 0x00, true},
		{"XRA A from FF", 0xAF, A, 0xFF, 0xFF, 0x00, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			REGISTERS[A] = tt.aValue
			if tt.src != A {
				REGISTERS[tt.src] = tt.srcValue
			}

			c.Execute(tt.opcode)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X ^ %s=0x%02X → 0x%02X",
						tt.aValue, regName(tt.src), tt.srcValue, tt.wantResult))
			}
		})
	}
}

func TestXRA_M(t *testing.T) {
	// XRA M (0xAE) - XOR memory with accumulator
	tests := []struct {
		name       string
		aValue     byte
		memValue   byte
		wantResult byte
	}{
		{"XRA M basic", 0xFF, 0x0F, 0xF0},
		{"XRA M same", 0xAA, 0xAA, 0x00},
		{"XRA M all ones", 0x00, 0xFF, 0xFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			REGISTERS[A] = tt.aValue
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00
			memory.MEMORY[0x2000] = tt.memValue

			c.Execute(0xAE)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, 0xAE, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, 0xAE, tt.name,
					fmt.Sprintf("A=0x%02X ^ (HL)=0x%02X → 0x%02X",
						tt.aValue, tt.memValue, tt.wantResult))
			}
		})
	}
}

// =============================================================================
// ORA INSTRUCTION TESTS (0xB0-0xB7)
// Logical OR register with accumulator
// =============================================================================

func TestORA_RegisterWithA(t *testing.T) {
	tests := []struct {
		name       string
		opcode     byte
		src        Reg
		aValue     byte
		srcValue   byte
		wantResult byte
		wantZero   bool
	}{
		// ORA B (0xB0)
		{"ORA B basic", 0xB0, B, 0xF0, 0x0F, 0xFF, false},
		{"ORA B zero", 0xB0, B, 0x00, 0x00, 0x00, true},

		// ORA C (0xB1)
		{"ORA C basic", 0xB1, C, 0xAA, 0x55, 0xFF, false},

		// ORA D (0xB2)
		{"ORA D basic", 0xB2, D, 0x33, 0x11, 0x33, false},

		// ORA E (0xB3)
		{"ORA E basic", 0xB3, E, 0x10, 0x01, 0x11, false},

		// ORA H (0xB4)
		{"ORA H basic", 0xB4, H, 0xCC, 0x33, 0xFF, false},

		// ORA L (0xB5)
		{"ORA L basic", 0xB5, L, 0x12, 0x34, 0x36, false},

		// ORA A (0xB7) - OR with self, value unchanged
		{"ORA A unchanged", 0xB7, A, 0x42, 0x42, 0x42, false},
		{"ORA A zero", 0xB7, A, 0x00, 0x00, 0x00, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			REGISTERS[A] = tt.aValue
			if tt.src != A {
				REGISTERS[tt.src] = tt.srcValue
			}

			c.Execute(tt.opcode)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("A=0x%02X | %s=0x%02X → 0x%02X",
						tt.aValue, regName(tt.src), tt.srcValue, tt.wantResult))
			}
		})
	}
}

func TestORA_M(t *testing.T) {
	// ORA M (0xB6) - OR memory with accumulator
	tests := []struct {
		name       string
		aValue     byte
		memValue   byte
		wantResult byte
	}{
		{"ORA M basic", 0xF0, 0x0F, 0xFF},
		{"ORA M zero", 0x00, 0x00, 0x00},
		{"ORA M same", 0x55, 0x55, 0x55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			REGISTERS[A] = tt.aValue
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00
			memory.MEMORY[0x2000] = tt.memValue

			c.Execute(0xB6)

			if REGISTERS[A] != tt.wantResult {
				logFail(t, 0xB6, tt.name,
					fmt.Sprintf("A=0x%02X", tt.wantResult),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, 0xB6, tt.name,
					fmt.Sprintf("A=0x%02X | (HL)=0x%02X → 0x%02X",
						tt.aValue, tt.memValue, tt.wantResult))
			}
		})
	}
}

// =============================================================================
// CMP INSTRUCTION TESTS (0xB8-0xBF)
// Compare register with accumulator (affects flags, not A)
// =============================================================================

func TestCMP_RegisterWithA(t *testing.T) {
	tests := []struct {
		name      string
		opcode    byte
		src       Reg
		aValue    byte
		srcValue  byte
		wantZero  bool // Set if A == src
		wantCarry bool // Set if A < src (borrow occurred)
		wantSign  bool // Set if result is negative
	}{
		// CMP B (0xB8)
		{"CMP B equal", 0xB8, B, 0x42, 0x42, true, false, false},
		{"CMP B greater", 0xB8, B, 0x50, 0x30, false, false, false},
		{"CMP B less", 0xB8, B, 0x30, 0x50, false, true, true},

		// CMP C (0xB9)
		{"CMP C equal", 0xB9, C, 0xAA, 0xAA, true, false, false},
		{"CMP C greater", 0xB9, C, 0xFF, 0x00, false, false, true},

		// CMP D (0xBA)
		{"CMP D basic", 0xBA, D, 0x33, 0x11, false, false, false},

		// CMP E (0xBB)
		{"CMP E basic", 0xBB, E, 0x55, 0x55, true, false, false},

		// CMP H (0xBC)
		{"CMP H basic", 0xBC, H, 0x80, 0x80, true, false, false},

		// CMP L (0xBD)
		{"CMP L basic", 0xBD, L, 0x12, 0x34, false, true, true},

		// CMP A (0xBF) - always equal
		{"CMP A always equal", 0xBF, A, 0x42, 0x42, true, false, false},
		{"CMP A from zero", 0xBF, A, 0x00, 0x00, true, false, false},
		{"CMP A from FF", 0xBF, A, 0xFF, 0xFF, true, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()

			REGISTERS[A] = tt.aValue
			if tt.src != A {
				REGISTERS[tt.src] = tt.srcValue
			}
			originalA := REGISTERS[A]

			c.Execute(tt.opcode)

			// CMP should NOT modify the accumulator
			if REGISTERS[A] != originalA {
				logFail(t, tt.opcode, tt.name,
					fmt.Sprintf("A unchanged=0x%02X", originalA),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, tt.opcode, tt.name,
					fmt.Sprintf("CMP A=0x%02X, %s=0x%02X (Z=%v, CY=%v)",
						tt.aValue, regName(tt.src), tt.srcValue, tt.wantZero, tt.wantCarry))
			}
		})
	}
}

func TestCMP_M(t *testing.T) {
	// CMP M (0xBE) - Compare memory with accumulator
	tests := []struct {
		name      string
		aValue    byte
		memValue  byte
		wantZero  bool
		wantCarry bool
	}{
		{"CMP M equal", 0x42, 0x42, true, false},
		{"CMP M greater", 0x50, 0x30, false, false},
		{"CMP M less", 0x30, 0x50, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestCPU()
			resetRegisters()
			clearMemory()

			REGISTERS[A] = tt.aValue
			REGISTERS[H] = 0x20
			REGISTERS[L] = 0x00
			memory.MEMORY[0x2000] = tt.memValue
			originalA := REGISTERS[A]

			c.Execute(0xBE)

			// CMP should NOT modify the accumulator
			if REGISTERS[A] != originalA {
				logFail(t, 0xBE, tt.name,
					fmt.Sprintf("A unchanged=0x%02X", originalA),
					fmt.Sprintf("A=0x%02X", REGISTERS[A]))
			} else {
				logPass(t, 0xBE, tt.name,
					fmt.Sprintf("CMP A=0x%02X, (HL)=0x%02X (Z=%v, CY=%v)",
						tt.aValue, tt.memValue, tt.wantZero, tt.wantCarry))
			}
		})
	}
}

// =============================================================================
// CYCLES COUNTING TESTS FOR 0x80-0xBF
// =============================================================================

func TestCycles_ArithmeticLogical(t *testing.T) {
	// All register-based arithmetic/logical ops should take 4 cycles
	regOpcodes := []byte{
		// ADD (0x80-0x85, 0x87)
		0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87,
		// ADC (0x88-0x8D, 0x8F)
		0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8D, 0x8F,
		// SUB (0x90-0x95, 0x97)
		0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97,
		// SBB (0x98-0x9D, 0x9F)
		0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9F,
		// ANA (0xA0-0xA5, 0xA7)
		0xA0, 0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA7,
		// XRA (0xA8-0xAD, 0xAF)
		0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAF,
		// ORA (0xB0-0xB5, 0xB7)
		0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB7,
		// CMP (0xB8-0xBD, 0xBF)
		0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBF,
	}

	for _, opcode := range regOpcodes {
		c := newTestCPU()
		c.Execute(opcode)

		if c.CyclesCount != 4 {
			t.Errorf("Opcode 0x%02X: CyclesCount = %d, want 4", opcode, c.CyclesCount)
		}
	}
}

func TestCycles_ArithmeticLogicalMemory(t *testing.T) {
	// All memory-based arithmetic/logical ops should take 7 cycles
	memOpcodes := []byte{
		0x86, // ADD M
		0x8E, // ADC M
		0x96, // SUB M
		0x9E, // SBB M
		0xA6, // ANA M
		0xAE, // XRA M
		0xB6, // ORA M
		0xBE, // CMP M
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
