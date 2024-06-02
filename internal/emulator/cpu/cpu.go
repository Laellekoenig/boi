package cpu

type word = uint16

type registerType uint8

type MemoryReader interface {
	ByteAt(word) byte
	WordAt(word) word
}

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

type flagType uint8

const (
	FlagZ flagType = iota
	FlagN
	FlagH
	FlagC
)

type Cpu struct {
	a byte
	f byte
	b byte
	c byte
	d byte
	e byte
	h byte
	l byte

	sp word
	pc word

	bus MemoryReader
}

func NewCpu(m MemoryReader) *Cpu {
	c := Cpu{
		bus: m,
	}
	c.reset()
	return &c
}

func (c *Cpu) reset() {
	c.a = 0x01
	c.f = 0xb0
	c.b = 0x00
	c.c = 0x13
	c.d = 0x00
	c.e = 0xd8
	c.h = 0x01
	c.l = 0x4d

	c.sp = 0xfffe
	c.pc = 0x0100
}

func (c *Cpu) ExecuteStep() {
	opcode := c.bus.ByteAt(c.pc)
	instruction := InstrucionFromByte(opcode)
	// TODO: handle timing
	// t := instruction.Timing()
	instruction.Execute(c)
}
