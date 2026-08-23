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

	var addr uint16 = (uint16(initialRegs[H]) << 8) | uint16(initialRegs[L])
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

func TestADC_X_WithCarry_NoOverflow(t *testing.T) {
	c := &CPU{}
	initialRegs := [7]byte{
		B: 0x5,
		C: 0x5,
		D: 0x5,
		E: 0x5,
		H: 0x5,
		L: 0x5,
		A: 0x5,
	}
	c.Init(initialRegs)

	var tests = []struct {
		name string
		reg  Reg
		want byte
	}{
		// With Carry Cases
		{"ADC_B", B, c.REGISTERS[A] + c.REGISTERS[B] + 1},
		{"ADC_C", C, c.REGISTERS[A] + c.REGISTERS[C] + 1},
		{"ADC_D", D, c.REGISTERS[A] + c.REGISTERS[D] + 1},
		{"ADC_E", E, c.REGISTERS[A] + c.REGISTERS[E] + 1},
		{"ADC_H", H, c.REGISTERS[A] + c.REGISTERS[H] + 1},
		{"ADC_L", L, c.REGISTERS[A] + c.REGISTERS[L] + 1},
		{"ADC_A", A, c.REGISTERS[A] + c.REGISTERS[A] + 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				c.Init(initialRegs)
			})
			c.SetFlag(CY)

			c.ADC_X(tt.reg)

			// TODO: look into putting this into a sub-test or some re-usable way. I suspect we'll have to 			// test flags like this quite a bit.
			if c.IsSet(CY) {
				t.Fatalf("CY flag should no be set")
			}
			if c.IsSet(Z) {
				t.Fatalf("Z flag should not be set")
			}
			if c.IsSet(S) {
				t.Fatalf("S flag should not be set")
			}
			if c.IsSet(P) {
				t.Fatalf("P flag should not be set")
			}
			if c.IsSet(AC) {
				t.Fatalf("AC flag should not be set")
			}

			if c.REGISTERS[A] != tt.want {
				t.Fatalf("got 0x%02X, want 0x%02X", c.REGISTERS[A], tt.want)
			}
		})
	}
}

func TestADC_X_WithCarryAndOverflow(t *testing.T) {
	c := &CPU{}
	initialRegs := [7]byte{
		B: 0x1,
		C: 0x1,
		D: 0x1,
		E: 0x1,
		H: 0x1,
		L: 0x1,
		A: 0xFE,
	}
	c.Init(initialRegs)

	var tests = []struct {
		name string
		reg  Reg
		want byte
	}{
		// With Carry Cases
		{"ADC_B", B, 0},
		{"ADC_C", C, 0},
		{"ADC_D", D, 0},
		{"ADC_E", E, 0},
		{"ADC_H", H, 0},
		{"ADC_L", L, 0},
		{"ADC_A", A, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				c.Init(initialRegs)
			})
			c.SetFlag(CY)

			c.ADC_X(tt.reg)

			if !c.IsSet(CY) {
				t.Fatalf("CY flag should be set")
			}
			if !c.IsSet(Z) {
				t.Fatalf("Z flag should be set")
			}
			if c.IsSet(S) {
				t.Fatalf("S flag should not be set")
			}
			if !c.IsSet(P) {
				t.Fatalf("P flag should be set")
			}
			if !c.IsSet(AC) {
				t.Fatalf("AC flag should be set")
			}

			if c.REGISTERS[A] != tt.want {
				t.Fatalf("got 0x%02X, want 0x%02X", c.REGISTERS[A], tt.want)
			}
		})
	}
}

func TestADC_X_WithoutCarryNoOverflow(t *testing.T) {
	c := &CPU{}
	initialRegs := [7]byte{
		B: 0x01,
		C: 0x01,
		D: 0x7F,
		E: 0x02,
		H: 0x03,
		L: 0x04,
		A: 0x01,
	}
	c.Init(initialRegs)

	var tests = []struct {
		name             string
		reg              Reg
		want             byte
		carryFlagWant    bool
		zeroFlagWant     bool
		signFlagWant     bool
		auxCarryFlagWant bool
		parityFlagWant   bool
	}{
		{
			name:             "ADC_B",
			reg:              B,
			want:             0x02,
			carryFlagWant:    false,
			zeroFlagWant:     false,
			signFlagWant:     false,
			auxCarryFlagWant: false,
			parityFlagWant:   false,
		},
		{
			name:             "ADC_C",
			reg:              C,
			want:             0x02,
			carryFlagWant:    false,
			zeroFlagWant:     false,
			signFlagWant:     false,
			auxCarryFlagWant: false,
			parityFlagWant:   false,
		},
		{
			name:             "ADC_D",
			reg:              D,
			want:             0x80,
			carryFlagWant:    false,
			zeroFlagWant:     false,
			signFlagWant:     true,
			auxCarryFlagWant: true,
			parityFlagWant:   false,
		},
		{
			name:             "ADC_E",
			reg:              E,
			want:             0x03,
			carryFlagWant:    false,
			zeroFlagWant:     false,
			signFlagWant:     false,
			auxCarryFlagWant: false,
			parityFlagWant:   true,
		},
		{
			name:             "ADC_H",
			reg:              H,
			want:             0x04,
			carryFlagWant:    false,
			zeroFlagWant:     false,
			signFlagWant:     false,
			auxCarryFlagWant: false,
			parityFlagWant:   false,
		},
		{
			name:             "ADC_L",
			reg:              L,
			want:             0x05,
			carryFlagWant:    false,
			zeroFlagWant:     false,
			signFlagWant:     false,
			auxCarryFlagWant: false,
			parityFlagWant:   true,
		},
		{
			name:             "ADC_A",
			reg:              A,
			want:             0x02,
			carryFlagWant:    false,
			zeroFlagWant:     false,
			signFlagWant:     false,
			auxCarryFlagWant: false,
			parityFlagWant:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				c.Init(initialRegs)
			})

			c.ADC_X(tt.reg)

			carryFlagGot := c.IsSet(CY)
			zeroFlagGot := c.IsSet(Z)
			signFlagGot := c.IsSet(S)
			parityFlagGot := c.IsSet(P)
			auxCarryFlagGot := c.IsSet(AC)

			if carryFlagGot != tt.carryFlagWant {
				t.Fatalf("A Reg [%b] | Got CY (Carry) flag as %t, want %t", c.REGISTERS[A], carryFlagGot, tt.carryFlagWant)
			}
			if zeroFlagGot != tt.zeroFlagWant {
				t.Fatalf("A Reg [%b] | Got Z (Zero) flag as %t, want %t", c.REGISTERS[A], zeroFlagGot, tt.zeroFlagWant)
			}
			if signFlagGot != tt.signFlagWant {
				t.Fatalf("A Reg [%b] | Got S (Sign) flag as %t, want %t", c.REGISTERS[A], signFlagGot, tt.signFlagWant)
			}
			if auxCarryFlagGot != tt.auxCarryFlagWant {
				t.Fatalf("A Reg [%b] | Got AC (Aux Carry) flag as %t, want %t", c.REGISTERS[A], auxCarryFlagGot, tt.auxCarryFlagWant)
			}
			if parityFlagGot != tt.parityFlagWant {
				t.Fatalf("A Reg [%b] | Got P (Parity) flag as %t, want %t", c.REGISTERS[A], parityFlagGot, tt.parityFlagWant)
			}

			if c.REGISTERS[A] != tt.want {
				t.Fatalf("got 0x%02X, want 0x%02X", c.REGISTERS[A], tt.want)
			}
		})
	}
}

/*
func TestADC_X_WithoutCarryOverflow(t *testing.T) {

}*/
