package cpu

type CPU struct {
	Regs [7]byte
	PC uint16
	CYCLES_TABLE [256]byte
	CyclesCount uint
}

