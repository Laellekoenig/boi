package cpu

type word = uint16

type MemoryReader interface {
	ByteAt(word) byte
	WordAt(word) word
	WriteByteAt(byte, word)
	WriteWordAt(word, word)
	Reset()
}

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

func (c *Cpu) currentByte() byte {
	b := c.bus.ByteAt(c.pc)
	c.pc += 1
	return b
}

func (c *Cpu) peakByte() byte {
	return c.bus.ByteAt(c.pc)
}

func (c *Cpu) currentWord() word {
	w := c.bus.WordAt(c.pc)
	c.pc += 2
	return w
}

func (c *Cpu) peakWord() word {
	return c.bus.WordAt(c.pc)
}

func (c *Cpu) dereferenceRegister(r doubleRegisterType) byte {
	return c.bus.ByteAt(c.readDoubleRegister(r))
}

func (c *Cpu) ExecuteStep() {
	opcode := c.currentByte()
	instruction := InstrucionFromByte(opcode)
	// TODO: handle timing
	_ = instruction.Execute(c)
}

func (c *Cpu) Reset() {
	c.reset()
	c.bus.Reset()
}
