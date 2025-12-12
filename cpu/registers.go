package cpu

var REGISTERS = [7]byte{         0,   0,   2,   0,   0,   0,   0}
var REGISTERS_NAMES = [7]string{"B", "C", "D", "E", "H", "L", "A"}
var FLAGS = [8]byte{0, 0, 0, 0, 0, 0, 0, 0}

type Reg byte
const (
	B Reg = iota
	C
	D
	E
	H
	L
	A
)

type Flag byte
const (
	F_C Flag = iota
	F_NOT_USED_1
	F_P
	F_NOT_USED_2
	F_A
	F_NOT_USED_3
	F_Z
	F_S
)
