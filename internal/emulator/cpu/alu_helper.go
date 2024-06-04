package cpu

func (c *Cpu) incRegs(r doubleRegisterType) {
	v := c.readDoubleRegister(r) + 1
	c.writeDoubleRegister(v, r)
}

func (c *Cpu) decRegs(r doubleRegisterType) {
	v := c.readDoubleRegister(r) - 1
	c.writeDoubleRegister(v, r)
}

func (c *Cpu) incReg(t registerType) {
	oldVal := c.readRegister(t)
	c.writeRegister(oldVal+1, t)

	c.setFlag(FlagZ, oldVal == 0xff)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, doesByteAddHalfCarry(oldVal, 1))
}

func (c *Cpu) decReg(t registerType) {
	oldVal := c.readRegister(t)
	c.writeRegister(oldVal-1, t)

	c.setFlag(FlagZ, oldVal == 1)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, doesByteSubHalfCarry(oldVal, 1))
}

func (c *Cpu) addRegsRegs(dst, src doubleRegisterType) {
	a := c.readDoubleRegister(dst)
	b := c.readDoubleRegister(src)
	c.writeDoubleRegister(a+b, dst)

	c.setFlag(FlagN, false)
	c.setFlag(FlagH, doesWordAddHalfCarry(a, b))
	c.setFlag(FlagC, doesWordAddCarry(a, b))
}

func (c *Cpu) daa() {
	a := c.readRegister(RegA)
	carry := c.checkFlag(FlagC)

	if !c.checkFlag(FlagN) {
		if c.checkFlag(FlagC) || a > 0x99 {
			a += 0x60
			c.setFlag(FlagC, true)
		}

		if c.checkFlag(FlagH) || (a&0x0f) > 0x09 {
			a += 0x06
		}

	} else {
		if c.checkFlag(FlagC) {
			a -= 0x60
		}

		if c.checkFlag(FlagH) {
			a -= 0x06
		}
	}
	c.writeRegister(a, RegA)

	c.setFlag(FlagZ, a == 0)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, carry || (!c.checkFlag(FlagN) && a > 0x99))
}

func (c *Cpu) cpl() {
	a := c.readRegister(RegA)
	a = ^a
	c.writeRegister(a, RegA)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, true)
}

func (c *Cpu) incDerefHL() {
	addr := c.readDoubleRegister(RegsHL)
	v := c.bus.ByteAt(addr)
	c.bus.WriteByteAt(v+1, addr)
	c.setFlag(FlagZ, v+1 == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, doesByteAddHalfCarry(v, 1))
}

func (c *Cpu) decDerefHL() {
	addr := c.readDoubleRegister(RegsHL)
	v := c.bus.ByteAt(addr)
	c.bus.WriteByteAt(v-1, addr)
	c.setFlag(FlagZ, v-1 == 0)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, doesByteSubHalfCarry(v, 1))
}

func (c *Cpu) addHLSP() {
	a := c.readDoubleRegister(RegsHL)
	b := c.sp
	c.writeDoubleRegister(a+b, RegsHL)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, doesWordAddHalfCarry(a, b))
	c.setFlag(FlagC, doesWordAddCarry(a, b))
}
