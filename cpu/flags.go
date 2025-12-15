package cpu

var FLAGS byte = 0

const (
	CY byte = 0x80 // Carry flag (bit 7)
	P  byte = 0x20 // Parity flag (bit 5)
	AC byte = 0x08 // Auxiliary Carry flag (bit 3)
	Z  byte = 0x02 // Zero flag (bit 1)
	S  byte = 0x01 // Sign flag (bit 0)
)

func SetFlag(masks byte) {
	FLAGS = FLAGS | masks
}

func ClearFlag(masks byte) {
	FLAGS = FLAGS &^ masks
}

func IsSet(mask byte) bool {
	return FLAGS&mask != 0
}
