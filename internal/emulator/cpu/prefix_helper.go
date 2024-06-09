package cpu

func (c *Cpu) rlca() uint8 {
	a := c.readRegister(RegA)
	carry := a&0x80 != 0
	c.writeRegister(a<<1|a&0x80>>7, RegA)

	c.setFlag(FlagZ, false)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, carry)

	return 4
}

func (c *Cpu) rrca() uint8 {
	a := c.readRegister(RegA)
	carry := a&1 != 0
	c.writeRegister(a>>1|a&1<<7, RegA)

	c.setFlag(FlagZ, false)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, carry)

	return 4
}

func (c *Cpu) rla() uint8 {
	a := c.readRegister(RegA)
	carry := a&0x80 != 0

	val := a << 1
	if c.checkFlag(FlagC) {
		val |= 1
	}
	c.writeRegister(val, RegA)

	c.setFlag(FlagZ, false)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, carry)

	return 4
}

func (c *Cpu) rra() uint8 {
	a := c.readRegister(RegA)
	carry := a&1 != 0

	val := a >> 1
	if c.checkFlag(FlagC) {
		val |= 0x80
	}
	c.writeRegister(val, RegA)

	c.setFlag(FlagZ, false)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, carry)

	return 4
}

func (c *Cpu) srlReg(r registerType) uint8 {
	val := c.readRegister(r)
	res := val >> 1
	c.writeRegister(res, r)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&1 > 0)

	return 8
}
func (c *Cpu) srlDerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val >> 1
	c.bus.WriteByteAt(res, addr)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&1 > 0)

	return 16
}

func (c *Cpu) rrReg(r registerType) uint8 {
	val := c.readRegister(r)
	res := val >> 1
	if c.checkFlag(FlagC) {
		res |= 0x80
	}
	c.writeRegister(res, r)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&1 > 0)

	return 8
}

func (c *Cpu) rrDerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val >> 1
	if c.checkFlag(FlagC) {
		res |= 0x80
	}
	c.bus.WriteByteAt(res, addr)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&1 > 0)

	return 16
}

func (c *Cpu) rlcReg(r registerType) uint8 {
	val := c.readRegister(r)
	res := val<<1 | val&0x80>>7
	c.writeRegister(res, r)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&0x80 > 0)

	return 8
}

func (c *Cpu) rlcDerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val<<1 | val&0x80>>7
	c.bus.WriteByteAt(res, addr)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&0x80 > 0)

	return 16
}

func (c *Cpu) rrcReg(r registerType) uint8 {
	val := c.readRegister(r)
	res := val>>1 | val&1<<7
	c.writeRegister(res, r)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&1 > 0)

	return 8
}

func (c *Cpu) rrcDerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val>>1 | val&1<<7
	c.bus.WriteByteAt(res, addr)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&1 > 0)

	return 16
}

func (c *Cpu) rlReg(r registerType) uint8 {
	val := c.readRegister(r)
	res := val << 1

	if c.checkFlag(FlagC) {
		res |= 1
	}

	c.writeRegister(res, r)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&0x80 > 0)

	return 8
}

func (c *Cpu) rlDerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val << 1

	if c.checkFlag(FlagC) {
		res |= 1
	}

	c.bus.WriteByteAt(res, addr)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&0x80 > 0)

	return 16
}

func (c *Cpu) slaReg(r registerType) uint8 {
	val := c.readRegister(r)
	res := val << 1
	c.writeRegister(res, r)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&0x80 > 0)

	return 8
}

func (c *Cpu) slaDerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val << 1
	c.bus.WriteByteAt(res, addr)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&0x80 > 0)

	return 16
}

func (c *Cpu) sraReg(r registerType) uint8 {
	val := c.readRegister(r)
	res := val>>1 | val&0x80
	c.writeRegister(res, r)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&1 > 0)

	return 8
}

func (c *Cpu) sraDerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val>>1 | val&0x80
	c.bus.WriteByteAt(res, addr)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, val&1 > 0)

	return 16
}

func (c *Cpu) swapReg(r registerType) uint8 {
	val := c.readRegister(r)
	upper := val & 0xf0
	lower := val & 0x0f
	res := (upper >> 4) | (lower << 4)
	c.writeRegister(res, r)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, false)

	return 8
}

func (c *Cpu) swapDerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	upper := val & 0xf0
	lower := val & 0x0f
	res := (upper >> 4) | (lower << 4)
	c.bus.WriteByteAt(res, addr)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, false)

	return 16
}

func (c *Cpu) bitReg(n uint8, r registerType) uint8 {
	val := c.readRegister(r)
	var mask byte
	mask = 1 << n

	c.setFlag(FlagZ, val&mask == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, true)

	return 8
}

func (c *Cpu) bitDerefHL(n uint8) uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	var mask byte
	mask = 1 << n

	c.setFlag(FlagZ, val&mask == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, true)

	return 12
}

func (c *Cpu) resReg(n uint8, r registerType) uint8 {
	val := c.readRegister(r)
	res := val & (^(1 << n))
	c.writeRegister(res, r)

	return 8
}

func (c *Cpu) resDerefHL(n uint8) uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val & (^(1 << n))
	c.bus.WriteByteAt(res, addr)

	return 16
}

func (c *Cpu) setReg(n uint8, r registerType) uint8 {
	val := c.readRegister(r)
	res := val | (1 << n)
	c.writeRegister(res, r)

	return 8
}

func (c *Cpu) setDerefHL(n uint8) uint8 {
	addr := c.readDoubleRegister(RegsHL)
	val := c.bus.ByteAt(addr)
	res := val | (1 << n)
	c.bus.WriteByteAt(res, addr)

	return 16
}
