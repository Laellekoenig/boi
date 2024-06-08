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
