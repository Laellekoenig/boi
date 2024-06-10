package cpu

type flagType uint8

const (
	FlagZ flagType = 1 << 7
	FlagN flagType = 1 << 6
	FlagH flagType = 1 << 5
	FlagC flagType = 1 << 4
)

func (c *Cpu) checkFlag(ft flagType) bool {
	return c.f&uint8(ft) > 0
}

func (c *Cpu) setFlag(ft flagType, f bool) {
	if f {
		c.f |= uint8(ft)
	} else {
		c.f &= ^uint8(ft)
	}
}
