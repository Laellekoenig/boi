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

func (c *Cpu) getRegister(t registerType) *byte {
	switch t {
	case RegA:
		return &c.a
	case RegF:
		return &c.f
	case RegB:
		return &c.b
	case RegC:
		return &c.c
	case RegD:
		return &c.d
	case RegE:
		return &c.e
	case RegH:
		return &c.h
	case RegL:
		return &c.l
	default:
		panic("invalid register type")
	}
}

func (c *Cpu) readRegister(t registerType) byte {
	return *c.getRegister(t)
}

func (c *Cpu) writeRegister(b byte, t registerType) {
	r := c.getRegister(t)
	*r = b
}

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

	if r == RegsAF {
		*bs[1] = byte(w & 0xf0)
	} else {
		*bs[1] = byte(w & 0xff)
	}
}
