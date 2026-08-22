package cpu

import (
	"8080/memory"
	"testing"
)

func TestADD_X(t *testing.T) {
	c := &CPU{}
	initialRegs := [7]byte{0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x0}
	c.Init(initialRegs)

	var tests = []struct {
		name string
		reg  Reg
		want byte
	}{
		{"ADD B", B, c.REGISTERS[B] + c.REGISTERS[A]},
		{"ADD C", C, c.REGISTERS[C] + c.REGISTERS[A]},
		{"ADD D", D, c.REGISTERS[D] + c.REGISTERS[A]},
		{"ADD E", E, c.REGISTERS[E] + c.REGISTERS[A]},
		{"ADD H", H, c.REGISTERS[H] + c.REGISTERS[A]},
		{"ADD L", L, c.REGISTERS[L] + c.REGISTERS[A]},
		{"ADD A", A, c.REGISTERS[A] + c.REGISTERS[A]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				c.Init(initialRegs)
			})

			c.ADD_X(tt.reg)
			if c.REGISTERS[A] != tt.want {
				t.Fatalf("got 0x%02X, want 0x%02X", c.REGISTERS[A], tt.want)
			}
		})
	}
}

func TestADD_M_X(t *testing.T) {
	c := &CPU{}
	initialRegs := [7]byte{
		B: 0xA,
		H: 0xFF,
		L: 0xFF,
	}
	c.Init(initialRegs)

	addr := (uint16(initialRegs[H]) << 8) | uint16(initialRegs[L])
	memory.MEMORY[addr] = 0x0A

	t.Cleanup(func() {
		memory.MEMORY[addr] = 0
	})

	c.ADD_M_X(B)

	const want byte = 0x14
	if got := c.REGISTERS[A]; got != want {
		t.Fatalf("got 0x%02X, want 0x%02X", memory.MEMORY[addr], want)
	}
}
