package cpu

type Instruction interface {
	String() string
	Execute(c *Cpu) uint8
}

type NotImplemented struct{}

func (n *NotImplemented) String() string       { return "Not implemented" }
func (n *NotImplemented) Execute(c *Cpu) uint8 { panic("instruction not implemented") }

type NOP struct{}

func (n *NOP) String() string { return "NOP" }
func (n *NOP) Execute(c *Cpu) uint8 {
	return 4
}

type JPWord struct{}

func (j *JPWord) String() string { return "JP u16" }
func (j *JPWord) Execute(c *Cpu) uint8 {
	addr := c.currentWord()
	c.jump(addr, None)
	return 16
}

type LdHlU16 struct{}

func (l *LdHlU16) String() string { return "LD HL, u16" }
func (l *LdHlU16) Execute(c *Cpu) uint8 {
	from := &readWord{c.currentWord()}
	to := &doubleRegister{[2]*byte{&c.h, &c.l}}
	loadWord(from, to)
	return 12
}

type LdBA struct{}

func (l *LdBA) String() string { return "LD B, A" }
func (l *LdBA) Execute(c *Cpu) uint8 {
	from := &readRegister{&c.a}
	to := &writeRegister{&c.b}
	loadByte(from, to)
	return 4
}

type LdDEU16 struct{}

func (l *LdDEU16) String() string { return "LD DE, u16" }
func (l *LdDEU16) Execute(c *Cpu) uint8 {
	from := &readWord{c.currentWord()}
	to := &doubleRegister{[2]*byte{&c.d, &c.e}}
	loadWord(from, to)
	return 12
}

type LdCU8 struct{}

func (l *LdCU8) String() string { return "LD C, u8" }
func (l *LdCU8) Execute(c *Cpu) uint8 {
	from := &readByte{c.currentByte()}
	to := &writeRegister{&c.c}
	loadByte(from, to)
	return 8
}

type LdAHLPlus struct{}

func (l *LdAHLPlus) String() string { return "LD, A, (HL+)" }
func (l *LdAHLPlus) Execute(c *Cpu) uint8 {
	b := c.dereferenceRegister(RegsHL)
	from := &readByte{b}
	to := &writeRegister{&c.a}
	loadByte(from, to)
	c.incDoubleRegister(RegsHL)
	return 8
}

func InstrucionFromByte(opcode byte) Instruction {
	switch opcode {
	case 0x00:
		return &NOP{}
	case 0x0e:
		return &LdCU8{}
	case 0x11:
		return &LdDEU16{}
	case 0x21:
		return &LdHlU16{}
	case 0x2a:
		return &LdAHLPlus{}
	case 0x47:
		return &LdBA{}
	case 0xc3:
		return &JPWord{}
	}

	return &NotImplemented{}
}
