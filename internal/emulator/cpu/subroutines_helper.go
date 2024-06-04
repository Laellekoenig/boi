package cpu

type jumpCondition uint8

const (
	None jumpCondition = iota
	NZ
	NC
	Z
	C
)

func (c *Cpu) jumpIm(condition jumpCondition) uint8 {
	switch condition {
	case None:
		// nothing
	case NZ:
		if c.checkFlag(FlagZ) {
			return 12
		}
	case NC:
		if c.checkFlag(FlagC) {
			return 12
		}
	case Z:
		if !c.checkFlag(FlagZ) {
			return 12
		}
	case C:
		if !c.checkFlag(FlagC) {
			return 12
		}
	default:
		panic("invalid condition")
	}

	c.pc = c.currentWord()

	return 16
}

func (c *Cpu) jumpRelativeIm(condition jumpCondition) uint8 {
	switch condition {
	case None:
		// nothing
	case NZ:
		if c.checkFlag(FlagZ) {
			return 8
		}
	case NC:
		if c.checkFlag(FlagC) {
			return 8
		}
	case Z:
		if !c.checkFlag(FlagZ) {
			return 8
		}
	case C:
		if !c.checkFlag(FlagC) {
			return 8
		}
	default:
		panic("invalid condition")
	}

	newPC := int32(c.pc) + int32(int8(c.currentByte()))
	c.pc = word(newPC)

	return 12
}

func (c *Cpu) ret(condition jumpCondition) uint8 {
	switch condition {
	case None:
		// nothing
	case NZ:
		if c.checkFlag(FlagZ) {
			return 8
		}
	case NC:
		if c.checkFlag(FlagC) {
			return 8
		}
	case Z:
		if !c.checkFlag(FlagZ) {
			return 8
		}
	case C:
		if !c.checkFlag(FlagC) {
			return 8
		}
	default:
		panic("invalid condition")
	}

	c.pc = c.popWord()

	return 20
}

func (c *Cpu) callIm(condition jumpCondition) uint8 {
	switch condition {
	case None:
		// nothing
	case NZ:
		if c.checkFlag(FlagZ) {
			return 12
		}
	case NC:
		if c.checkFlag(FlagC) {
			return 12
		}
	case Z:
		if !c.checkFlag(FlagZ) {
			return 12
		}
	case C:
		if !c.checkFlag(FlagC) {
			return 12
		}
	default:
		panic("invalid condition")
	}

	c.pushWord(c.pc + 3)
	c.pc = c.currentWord()

	return 24
}

func (c *Cpu) rst(addr word) uint8 {
	c.pushWord(c.pc)
	c.pc = addr
	return 16
}

func (c *Cpu) reti() uint8 {
	c.pc = c.popWord()
	c.ime = true
	return 16
}

func (c *Cpu) jpHL() uint8 {
	c.pc = c.readDoubleRegister(RegsHL)
	return 4
}