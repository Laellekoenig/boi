package cpu

type jumpCondition uint8

const (
	None jumpCondition = iota
	NZ
	NC
	Z
	C
)

func (c *Cpu) jump(address word, condition jumpCondition) uint8 {
	switch condition {
	case None:
		// nothing
	case NZ:
		if c.CheckFlag(FlagZ) {
			return 12
		}
	case NC:
		if c.CheckFlag(FlagC) {
			return 12
		}
	case Z:
		if !c.CheckFlag(FlagZ) {
			return 12
		}
	case C:
		if !c.CheckFlag(FlagC) {
			return 12
		}
	default:
		panic("invalid condition")
	}

	c.pc = address

	return 16
}
