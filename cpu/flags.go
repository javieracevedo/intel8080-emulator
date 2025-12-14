package cpu

var FLAGS = [8]byte{0, 0, 0, 0, 0, 0, 0, 0}

type Flag byte

const (
	CY Flag = iota
	F_NOT_USED_1
	P
	F_NOT_USED_2
	AC
	F_NOT_USED_3
	Z
	S
)
