package cpu

var REGISTERS = [7]byte{0, 0, 0, 0, 0, 0, 0}
var REGISTERS_NAMES = [7]string{"B", "C", "D", "E", "H", "L", "A"}

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
