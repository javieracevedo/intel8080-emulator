package cpu

import (
	"8080/memory"
	"testing"
)

func TestMOV(t *testing.T) {
	c := &CPU{}
	initialRegs := [7]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7}
	c.Init(initialRegs)

	var tests = []struct {
		name     string
		regLeft  Reg
		regRight Reg
		want     byte
	}{
		// MOV B,X
		{"MOV B,B", B, B, c.REGISTERS[B]},
		{"MOV B,C", B, C, c.REGISTERS[C]},
		{"MOV B,D", B, D, c.REGISTERS[D]},
		{"MOV B,E", B, E, c.REGISTERS[E]},
		{"MOV B,H", B, H, c.REGISTERS[H]},
		{"MOV B,L", B, L, c.REGISTERS[L]},
		{"MOV B,A", B, A, c.REGISTERS[A]},
		// MOV C,X
		{"MOV C,B", C, B, c.REGISTERS[B]},
		{"MOV C,C", C, C, c.REGISTERS[C]},
		{"MOV C,D", C, D, c.REGISTERS[D]},
		{"MOV C,E", C, E, c.REGISTERS[E]},
		{"MOV C,H", C, H, c.REGISTERS[H]},
		{"MOV C,L", C, L, c.REGISTERS[L]},
		{"MOV C,A", C, A, c.REGISTERS[A]},
		// MOV D,X
		{"MOV D,B", D, B, c.REGISTERS[B]},
		{"MOV D,C", D, C, c.REGISTERS[C]},
		{"MOV D,D", D, D, c.REGISTERS[D]},
		{"MOV D,E", D, E, c.REGISTERS[E]},
		{"MOV D,H", D, H, c.REGISTERS[H]},
		{"MOV D,L", D, L, c.REGISTERS[L]},
		{"MOV D,A", D, A, c.REGISTERS[A]},
		// MOV E,X
		{"MOV E,B", E, B, c.REGISTERS[B]},
		{"MOV E,C", E, C, c.REGISTERS[C]},
		{"MOV E,D", E, D, c.REGISTERS[D]},
		{"MOV E,E", E, E, c.REGISTERS[E]},
		{"MOV E,H", E, H, c.REGISTERS[H]},
		{"MOV E,L", E, L, c.REGISTERS[L]},
		{"MOV E,A", E, A, c.REGISTERS[A]},
		// MOV H,X
		{"MOV H,B", H, B, c.REGISTERS[B]},
		{"MOV H,C", H, C, c.REGISTERS[C]},
		{"MOV H,D", H, D, c.REGISTERS[D]},
		{"MOV H,E", H, E, c.REGISTERS[E]},
		{"MOV H,H", H, H, c.REGISTERS[H]},
		{"MOV H,L", H, L, c.REGISTERS[L]},
		{"MOV H,A", H, A, c.REGISTERS[A]},
		// MOV L,X
		{"MOV L,B", L, B, c.REGISTERS[B]},
		{"MOV L,C", L, C, c.REGISTERS[C]},
		{"MOV L,D", L, D, c.REGISTERS[D]},
		{"MOV L,E", L, E, c.REGISTERS[E]},
		{"MOV L,H", L, H, c.REGISTERS[H]},
		{"MOV L,L", L, L, c.REGISTERS[L]},
		{"MOV L,A", L, A, c.REGISTERS[A]},
		// MOV A,X
		{"MOV A,B", A, B, c.REGISTERS[B]},
		{"MOV A,C", A, C, c.REGISTERS[C]},
		{"MOV A,D", A, D, c.REGISTERS[D]},
		{"MOV A,E", A, E, c.REGISTERS[E]},
		{"MOV A,H", A, H, c.REGISTERS[H]},
		{"MOV A,L", A, L, c.REGISTERS[L]},
		{"MOV A,A", A, A, c.REGISTERS[A]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Init(initialRegs)
			t.Cleanup(func() {
				c.Init()
			})

			c.MOV(tt.regLeft, tt.regRight)
			if c.REGISTERS[tt.regLeft] != byte(tt.want) {
				t.Fatalf("got 0x%02X, want 0x%02X", c.REGISTERS[tt.regLeft], tt.want)
			}
		})
	}
}

func TestMOV_X_M(t *testing.T) {
	c := &CPU{}
	c.Init()
	c.REGISTERS[H] = 0x1
	c.REGISTERS[L] = 0x1
	addr := uint16(0x0101)
	memory.MEMORY[addr] = 0x2

	t.Cleanup(func() {
		memory.MEMORY = [64 * 1014]byte{}
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
			c.Init([7]byte{0, 0, 0, 0, 0x1, 0x1, 0})
			t.Cleanup(func() {
				c.Init()
			})
			c.MOV_X_M(tt.reg)
			if c.REGISTERS[tt.reg] != 0x2 {
				t.Fatalf("got 0x%02X, want 0x%02X", c.REGISTERS[tt.reg], 0x2)
			}
		})
	}
}
