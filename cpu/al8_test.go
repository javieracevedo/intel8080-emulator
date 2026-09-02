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

func TestADC_X(t *testing.T) {
	c := &CPU{}
	withoutCarryRegs := [7]byte{
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
	withCarryAndNoOverflowRegs := [7]byte{
		B: 0x01,
		C: 0x02,
		D: 0x03,
		E: 0x04,
		H: 0x05,
		L: 0x06,
		A: 0x10,
	}

	var tests = []adcTestCase{
		{
			name:        "WithoutCarry/ADC_B",
			initialRegs: withoutCarryRegs,
			reg:         B,
			want:        0x02,
			flags:       0,
		},
		{
			name:        "WithoutCarry/ADC_C",
			initialRegs: withoutCarryRegs,
			reg:         C,
			want:        0x02,
			flags:       0,
		},
		{
			name:        "WithoutCarry/ADC_D",
			initialRegs: withoutCarryRegs,
			reg:         D,
			want:        0x80,
			flags:       S | AC,
		},
		{
			name:        "WithoutCarry/ADC_E",
			initialRegs: withoutCarryRegs,
			reg:         E,
			want:        0x03,
			flags:       P,
		},
		{
			name:        "WithoutCarry/ADC_H",
			initialRegs: withoutCarryRegs,
			reg:         H,
			want:        0x04,
			flags:       0x0,
		},
		{
			name:        "WithoutCarry/ADC_L",
			initialRegs: withoutCarryRegs,
			reg:         L,
			want:        0x05,
			flags:       P,
		},
		{
			name:        "WithoutCarry/ADC_A",
			initialRegs: withoutCarryRegs,
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
		{
			name:          "WithCarryAndNoOverflow/ADC_B",
			initialRegs:   withCarryAndNoOverflowRegs,
			isCarryPreset: true,
			reg:           B,
			want:          0x12,
			flags:         P,
		},
		{
			name:          "WithCarryAndNoOverflow/ADC_C",
			initialRegs:   withCarryAndNoOverflowRegs,
			isCarryPreset: true,
			reg:           C,
			want:          0x13,
			flags:         0x0,
		},
		{
			name:          "WithCarryAndNoOverflow/ADC_D",
			initialRegs:   withCarryAndNoOverflowRegs,
			isCarryPreset: true,
			reg:           D,
			want:          0x14,
			flags:         P,
		},
		{
			name:          "WithCarryAndNoOverflow/ADC_E",
			initialRegs:   withCarryAndNoOverflowRegs,
			isCarryPreset: true,
			reg:           E,
			want:          0x15,
			flags:         0x0,
		},
		{
			name:          "WithCarryAndNoOverflow/ADC_H",
			initialRegs:   withCarryAndNoOverflowRegs,
			isCarryPreset: true,
			reg:           H,
			want:          0x16,
			flags:         0x0,
		},
		{
			name:          "WithCarryAndNoOverflow/ADC_L",
			initialRegs:   withCarryAndNoOverflowRegs,
			isCarryPreset: true,
			reg:           L,
			want:          0x17,
			flags:         P,
		},
		{
			name:          "WithCarryAndNoOverflow/ADC_A",
			initialRegs:   withCarryAndNoOverflowRegs,
			isCarryPreset: true,
			reg:           A,
			want:          0x21,
			flags:         P,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Init(tt.initialRegs)
			c.Flags = 0
			if tt.isCarryPreset {
				c.SetFlag(CY)
			}

			c.ADC_X(tt.reg)

			assertFlags(t, c, tt.flags, c.Flags)

			if c.REGISTERS[A] != tt.want {
				t.Errorf("got 0x%02X, want 0x%02X", c.REGISTERS[A], tt.want)
			}
		})
	}
}

func assertFlags(t *testing.T, c *CPU, wantFlags byte, gotFlags byte) {
	t.Helper()

	flags := []struct {
		mask byte
		name string
	}{
		{CY, "CY (Carry)"},
		{Z, "Z (Zero)"},
		{S, "S (Sign)"},
		{AC, "AC (Aux Carry)"},
		{P, "P (Parity)"},
	}

	for _, f := range flags {
		if got, want := gotFlags&f.mask, wantFlags&f.mask; got != want {
			t.Errorf("A Reg [%08b] | Got %s flag as %b, want %b", c.REGISTERS[A], f.name, got, want)
		}
	}
}
