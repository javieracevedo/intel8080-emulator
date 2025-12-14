package cpu

import (
	"testing"
)

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
