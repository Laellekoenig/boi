package cpu

type registerType uint8

const (
	RegA registerType = iota
	RegF
	RegB
	RegC
	RegD
	RegE
	RegH
	RegL
)

type doubleRegisterType uint8

const (
	RegsAF doubleRegisterType = iota
	RegsBC
	RegsDE
	RegsHL
)

func (c *Cpu) getDoubleRegister(r doubleRegisterType) [2]*byte {
	var a, b *byte

	switch r {
	case RegsAF:
		a = &c.a
		b = &c.f
	case RegsBC:
		a = &c.b
		b = &c.c
	case RegsDE:
		a = &c.d
		b = &c.e
	case RegsHL:
		a = &c.h
		b = &c.l
	default:
		panic("invalid double register type")
	}

	return [2]*byte{a, b}
}

func (c *Cpu) readDoubleRegister(r doubleRegisterType) word {
	bs := c.getDoubleRegister(r)
	return word(*bs[0])<<8 | word(*bs[1])
}

func (c *Cpu) writeDoubleRegister(w word, r doubleRegisterType) {
	bs := c.getDoubleRegister(r)
	*bs[0] = byte(w >> 8)
	*bs[1] = byte(w & 0xff)
}

func (c *Cpu) incDoubleRegister(r doubleRegisterType) {
	v := c.readDoubleRegister(r) + 1
	c.writeDoubleRegister(v, r)
}
