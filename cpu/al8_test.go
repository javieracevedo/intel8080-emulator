package cpu

import (
	"8080/memory"
	"testing"
)

type adcTestCase struct {
	name          string
	initialRegs   [7]byte
	isCarryPreset bool
	reg           Reg
	want          byte
	flags         byte
}

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

// Test WithoutCarry and Overflow
// Test WithoutCarry and No Overflow
func TestADC_X(t *testing.T) {
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
	withCarryAndOverflowRegs := [7]byte{
		B: 0x01,
		C: 0x01,
		D: 0x01,
		E: 0x01,
		H: 0x01,
		L: 0x01,
		A: 0xFE,
	}

	var tests = []adcTestCase{
		{
			name:        "WithoutCarry/ADC_B",
			initialRegs: initialRegs,
			reg:         B,
			want:        0x02,
			flags:       0,
		},
		{
			name:        "WithoutCarry/ADC_C",
			initialRegs: initialRegs,
			reg:         C,
			want:        0x02,
			flags:       0,
		},
		{
			name:        "WithoutCarry/ADC_D",
			initialRegs: initialRegs,
			reg:         D,
			want:        0x80,
			flags:       S | AC,
		},
		{
			name:        "WithoutCarry/ADC_E",
			initialRegs: initialRegs,
			reg:         E,
			want:        0x03,
			flags:       P,
		},
		{
			name:        "WithoutCarry/ADC_H",
			initialRegs: initialRegs,
			reg:         H,
			want:        0x04,
			flags:       0x0,
		},
		{
			name:        "WithoutCarry/ADC_L",
			initialRegs: initialRegs,
			reg:         L,
			want:        0x05,
			flags:       P,
		},
		{
			name:        "WithoutCarry/ADC_A",
			initialRegs: initialRegs,
			reg:         A,
			want:        0x02,
			flags:       0x0,
		},
		{
			name:          "WithCarryAndOverflow/ADC_B",
			initialRegs:   withCarryAndOverflowRegs,
			isCarryPreset: true,
			reg:           B,
			want:          0x00,
			flags:         CY | Z | P | AC,
		},
		{
			name:          "WithCarryAndOverflow/ADC_C",
			initialRegs:   withCarryAndOverflowRegs,
			isCarryPreset: true,
			reg:           C,
			want:          0x00,
			flags:         CY | Z | P | AC,
		},
		{
			name:          "WithCarryAndOverflow/ADC_D",
			initialRegs:   withCarryAndOverflowRegs,
			isCarryPreset: true,
			reg:           D,
			want:          0x00,
			flags:         CY | Z | P | AC,
		},
		{
			name:          "WithCarryAndOverflow/ADC_E",
			initialRegs:   withCarryAndOverflowRegs,
			isCarryPreset: true,
			reg:           E,
			want:          0x00,
			flags:         CY | Z | P | AC,
		},
		{
			name:          "WithCarryAndOverflow/ADC_H",
			initialRegs:   withCarryAndOverflowRegs,
			isCarryPreset: true,
			reg:           H,
			want:          0x00,
			flags:         CY | Z | P | AC,
		},
		{
			name:          "WithCarryAndOverflow/ADC_L",
			initialRegs:   withCarryAndOverflowRegs,
			isCarryPreset: true,
			reg:           L,
			want:          0x00,
			flags:         CY | Z | P | AC,
		},
		{
			name:          "WithCarryAndOverflow/ADC_A",
			initialRegs:   withCarryAndOverflowRegs,
			isCarryPreset: true,
			reg:           A,
			want:          0xFD,
			flags:         CY | S | AC,
		},
	}

	// CY | P
	// 1010 0000

	// FLAGS&CY = 1000 0000
	// FLAGS&P  = 0010 0000

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Init(tt.initialRegs)
			c.Flags = 0
			if tt.isCarryPreset {
				c.SetFlag(CY)
			}

			c.ADC_X(tt.reg)

			/*
				carryFlagGot := c.IsSet(CY)
				zeroFlagGot := c.IsSet(Z)
				signFlagGot := c.IsSet(S)
				parityFlagGot := c.IsSet(P)
				auxCarryFlagGot := c.IsSet(AC)
			*/
			assertFlags(tt.flags, c.Flags, t, c)

			/*if carryFlagGot != tt.carryFlagWant {
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
			*/

			if c.REGISTERS[A] != tt.want {
				t.Fatalf("got 0x%02X, want 0x%02X", c.REGISTERS[A], tt.want)
			}
		})
	}
}

func assertFlags(wantFlags byte, gotFlags byte, t *testing.T, c *CPU) {
	carryFlagGot := gotFlags & CY
	zeroFlagGot := gotFlags & Z
	signFlagGot := gotFlags & S
	parityFlagGot := gotFlags & P
	auxCarryFlagGot := gotFlags & AC

	carryFlagWant := wantFlags & CY
	zeroFlagWant := wantFlags & Z
	signFlagWant := wantFlags & S
	parityFlagWant := wantFlags & P
	auxCarryFlagWant := wantFlags & AC

	if wantFlags&CY != gotFlags&CY {
		t.Fatalf("A Reg [%b] | Got CY (Carry) flag as %b, want %b", c.REGISTERS[A], carryFlagGot, carryFlagWant)
	}
	if wantFlags&Z != gotFlags&Z {
		t.Fatalf("A Reg [%b] | Got Z (Zero) flag as %b, want %b", c.REGISTERS[A], zeroFlagGot, zeroFlagWant)
	}
	if wantFlags&S != gotFlags&S {
		t.Fatalf("A Reg [%b] | Got S (Sign) flag as %b, want %b", c.REGISTERS[A], signFlagGot, signFlagWant)
	}
	if wantFlags&AC != gotFlags&AC {
		t.Fatalf("A Reg [%b] | Got AC (Aux Carry) flag as %b, want %b", c.REGISTERS[A], auxCarryFlagGot, auxCarryFlagWant)
	}
	if wantFlags&P != gotFlags&P {
		t.Fatalf("A Reg [%b] | Got P (Parity) flag as %b, want %b", c.REGISTERS[A], parityFlagGot, parityFlagWant)
	}
}

/*
func TestADC_X_WithoutCarryOverflow(t *testing.T) {

}*/
