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
