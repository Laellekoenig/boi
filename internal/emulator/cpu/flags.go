package cpu

type flagType uint8

const (
	FlagZ flagType = 7
	FlagN flagType = 6
	FlagH flagType = 5
	FlagC flagType = 4
)

func maskBit(i uint8) uint8 {
	return 1 << i
}

func (c *Cpu) checkFlag(ft flagType) bool {
	return c.f&maskBit(uint8(ft)) > 0
}

func (c *Cpu) setFlag(ft flagType, f bool) {
	if f {
		c.f |= maskBit(uint8(ft))
	} else {
		c.f &= ^maskBit(uint8(ft))
	}
}
