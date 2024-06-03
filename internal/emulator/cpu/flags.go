package cpu

type flagType uint8

const (
	FlagZ flagType = iota
	FlagN
	FlagH
	FlagC
)

func maskBit(i uint8) uint8 {
	return 1 << i
}

func (c *Cpu) checkFlag(ft flagType) bool {
	switch ft {
	case FlagZ:
		return c.f&maskBit(7) > 0
	case FlagN:
		return c.f&maskBit(6) > 0
	case FlagH:
		return c.f&maskBit(5) > 0
	case FlagC:
		return c.f&maskBit(4) > 0
	}
	panic("Invalid flag")
}

func (c *Cpu) setFlag(ft flagType, f bool) {
	switch ft {
	case FlagZ:
		if f {
			c.f |= maskBit(7)
		} else {
			c.f &= ^maskBit(7)
		}
	case FlagN:
		if f {
			c.f |= maskBit(6)
		} else {
			c.f &= ^maskBit(6)
		}
	case FlagH:
		if f {
			c.f |= maskBit(5)
		} else {
			c.f &= ^maskBit(5)
		}
	case FlagC:
		if f {
			c.f |= maskBit(4)
		} else {
			c.f &= ^maskBit(4)
		}
	default:
		panic("Invalid flag")
	}
}
