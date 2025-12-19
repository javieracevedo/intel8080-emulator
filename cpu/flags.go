package cpu

const (
	CY byte = 0x80 // Carry flag (bit 7)
	P  byte = 0x20 // Parity flag (bit 5)
	AC byte = 0x08 // Auxiliary Carry flag (bit 3)
	Z  byte = 0x02 // Zero flag (bit 1)
	S  byte = 0x01 // Sign flag (bit 0)
)

func (c *CPU) SetFlag(masks byte) {
	c.Flags = c.Flags | masks
}

func (c *CPU) ClearFlag(masks byte) {
	c.Flags = c.Flags &^ masks
}

func (c *CPU) IsSet(mask byte) bool {
	return c.Flags&mask != 0
}
