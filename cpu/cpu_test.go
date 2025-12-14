package cpu

import (
	"8080/memory"
	"testing"
)

func resetRegisters() {
	REGISTERS = [7]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7}
}

func resetMemory() {
	memory.MEMORY = [64 * 1014]byte{}
}

func TestMOV(t *testing.T) {
	resetRegisters()

	t.Cleanup(func() {
		resetRegisters()
	})

	var tests = []struct {
		name     string
		regLeft  Reg
		regRight Reg
		want     byte
	}{
		// MOV B,X
		{"MOV B,B", B, B, REGISTERS[B]},
		{"MOV B,C", B, C, REGISTERS[C]},
		{"MOV B,D", B, D, REGISTERS[D]},
		{"MOV B,E", B, E, REGISTERS[E]},
		{"MOV B,H", B, H, REGISTERS[H]},
		{"MOV B,L", B, L, REGISTERS[L]},
		{"MOV B,A", B, A, REGISTERS[A]},
		// MOV C,X
		{"MOV C,B", C, B, REGISTERS[B]},
		{"MOV C,C", C, C, REGISTERS[C]},
		{"MOV C,D", C, D, REGISTERS[D]},
		{"MOV C,E", C, E, REGISTERS[E]},
		{"MOV C,H", C, H, REGISTERS[H]},
		{"MOV C,L", C, L, REGISTERS[L]},
		{"MOV C,A", C, A, REGISTERS[A]},
		// MOV D,X
		{"MOV D,B", D, B, REGISTERS[B]},
		{"MOV D,C", D, C, REGISTERS[C]},
		{"MOV D,D", D, D, REGISTERS[D]},
		{"MOV D,E", D, E, REGISTERS[E]},
		{"MOV D,H", D, H, REGISTERS[H]},
		{"MOV D,L", D, L, REGISTERS[L]},
		{"MOV D,A", D, A, REGISTERS[A]},
		// MOV E,X
		{"MOV E,B", E, B, REGISTERS[B]},
		{"MOV E,C", E, C, REGISTERS[C]},
		{"MOV E,D", E, D, REGISTERS[D]},
		{"MOV E,E", E, E, REGISTERS[E]},
		{"MOV E,H", E, H, REGISTERS[H]},
		{"MOV E,L", E, L, REGISTERS[L]},
		{"MOV E,A", E, A, REGISTERS[A]},
		// MOV H,X
		{"MOV H,B", H, B, REGISTERS[B]},
		{"MOV H,C", H, C, REGISTERS[C]},
		{"MOV H,D", H, D, REGISTERS[D]},
		{"MOV H,E", H, E, REGISTERS[E]},
		{"MOV H,H", H, H, REGISTERS[H]},
		{"MOV H,L", H, L, REGISTERS[L]},
		{"MOV H,A", H, A, REGISTERS[A]},
		// MOV L,X
		{"MOV L,B", L, B, REGISTERS[B]},
		{"MOV L,C", L, C, REGISTERS[C]},
		{"MOV L,D", L, D, REGISTERS[D]},
		{"MOV L,E", L, E, REGISTERS[E]},
		{"MOV L,H", L, H, REGISTERS[H]},
		{"MOV L,L", L, L, REGISTERS[L]},
		{"MOV L,A", L, A, REGISTERS[A]},
		// MOV A,X
		{"MOV A,B", A, B, REGISTERS[B]},
		{"MOV A,C", A, C, REGISTERS[C]},
		{"MOV A,D", A, D, REGISTERS[D]},
		{"MOV A,E", A, E, REGISTERS[E]},
		{"MOV A,H", A, H, REGISTERS[H]},
		{"MOV A,L", A, L, REGISTERS[L]},
		{"MOV A,A", A, A, REGISTERS[A]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetRegisters()
			MOV(tt.regLeft, tt.regRight)
			if REGISTERS[tt.regLeft] != byte(tt.want) {
				t.Fatalf("got 0x%02X, want 0x%02X", REGISTERS[tt.regLeft], tt.want)
			}
		})
	}
}

func TestMOV_X_M(t *testing.T) {
	REGISTERS[H] = 0x1
	REGISTERS[L] = 0x1
	addr := 0x0101
	memory.MEMORY[addr] = 0x2

	t.Cleanup(func() {
		resetRegisters()
		resetMemory()
	})

	var tests = []struct {
		name string
		reg  Reg
		want byte
	}{
		{"MOV B,M", B, memory.MEMORY[0x0101]},
		{"MOV C,M", C, memory.MEMORY[0x0101]},
		{"MOV D,M", D, memory.MEMORY[0x0101]},
		{"MOV E,M", E, memory.MEMORY[0x0101]},
		{"MOV H,M", H, memory.MEMORY[0x0101]},
		{"MOV L,M", L, memory.MEMORY[0x0101]},
		{"MOV A,M", A, memory.MEMORY[0x0101]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MOV_X_M(tt.reg)
			if memory.MEMORY[0x0101] != 0x2 {
				t.Fatalf("got 0x%02X, want 0x%02X", memory.MEMORY[0x0101], 0x2)
			}
		})
	}
}

func TestADD_X(t *testing.T) {
	REGISTERS = [7]byte{0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x0}

	t.Cleanup(func() {
		resetRegisters()
		resetMemory()
	})

	var tests = []struct {
		name string
		reg  Reg
		want byte
	}{
		{"ADD B", B, REGISTERS[B] + REGISTERS[A]},
		{"ADD C", C, REGISTERS[C] + REGISTERS[A]},
		{"ADD D", D, REGISTERS[D] + REGISTERS[A]},
		{"ADD E", E, REGISTERS[E] + REGISTERS[A]},
		{"ADD H", H, REGISTERS[H] + REGISTERS[A]},
		{"ADD L", L, REGISTERS[L] + REGISTERS[A]},
		{"ADD A", A, REGISTERS[A] + REGISTERS[A]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			REGISTERS = [7]byte{0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x0}

			ADD_X(tt.reg)
			if REGISTERS[A] != tt.want {
				t.Fatalf("got 0x%02X, want 0x%02X", REGISTERS[A], tt.want)
			}
		})
	}
}
