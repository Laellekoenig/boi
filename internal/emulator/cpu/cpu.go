package cpu

import "fmt"

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

	ime bool

	bus MemoryReader

	PastOps []string
	counter int64
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
	opcode := c.peakByte()
	instruction := InstrucionFromByte(opcode)

	c.PastOps = append(c.PastOps, fmt.Sprintf("%d | %s: PC=%04x SP=%04x A=%02x F=%02x B=%02x C=%02x D=%02x E=%02x H=%02x L=%02x", c.counter, instruction.String(), c.pc, c.sp, c.a, c.f, c.b, c.c, c.d, c.e, c.h, c.l))
	c.pc += 1
	c.counter += 1

	// TODO: handle timing
	_ = instruction.Execute(c)
}

func (c *Cpu) ContinueUnimpl() {
	for {
		opcode := c.peakByte()
		instruction := InstrucionFromByte(opcode)

		if instruction.String() == "Not implemented" {
			break
		}

		c.PastOps = append(c.PastOps, fmt.Sprintf("%d | %s: PC=%04x SP=%04x A=%02x F=%02x B=%02x C=%02x D=%02x E=%02x H=%02x L=%02x", c.counter, instruction.String(), c.pc, c.sp, c.a, c.f, c.b, c.c, c.d, c.e, c.h, c.l))
		c.pc += 1
		c.counter += 1

		// TODO: handle timing
		_ = instruction.Execute(c)
	}
}

func (c *Cpu) ContinueUntil(bp word) {
	for c.pc != bp {
		c.ExecuteStep()
	}
}

func (c *Cpu) ContinueUntilFlagChange() {
	f := c.f
	for f == c.f {
		c.ExecuteStep()
	}
}

func (c *Cpu) Reset() {
	c.reset()
	c.bus.Reset()
}

func (c *Cpu) popWord() word {
	w := c.bus.WordAt(c.sp)
	c.sp += 2
	return w
}

func (c *Cpu) pushWord(w word) {
	c.sp -= 2
	c.bus.WriteWordAt(w, c.sp)
}
