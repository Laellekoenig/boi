package cpu

func maskBit(i uint8) uint8 {
	return 1 << i
}

func (c *Cpu) CheckFlag(ft flagType) bool {
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
