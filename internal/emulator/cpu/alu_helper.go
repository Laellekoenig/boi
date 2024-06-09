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

func (c *Cpu) addRegReg(a, b registerType) uint8 {
	aVal := c.readRegister(a)
	bVal := c.readRegister(b)
	res := aVal + bVal
	c.writeRegister(res, a)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, doesByteAddHalfCarry(aVal, bVal))
	c.setFlag(FlagC, doesByteAddCarry(aVal, bVal))

	return 4
}

func (c *Cpu) daa() {
	val := c.readRegister(RegA)

	var correction byte
	carry := false

	if c.checkFlag(FlagH) || (!c.checkFlag(FlagN) && val&0x0f > 9) {
		correction |= 0x06
	}

	if c.checkFlag(FlagC) || (!c.checkFlag(FlagN) && val > 0x99) {
		correction |= 0x60
		carry = true
	}

	if c.checkFlag(FlagN) {
		val -= correction
	} else {
		val += correction
	}

	c.writeRegister(val, RegA)

	c.setFlag(FlagZ, val == 0)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, carry)
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

func (c *Cpu) orRegReg(a, b registerType) uint8 {
	res := c.readRegister(a) | c.readRegister(b)
	c.writeRegister(res, a)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, false)

	return 4
}

func (c *Cpu) orADerefHL() uint8 {
	res := c.readRegister(RegA) | c.bus.ByteAt(c.readDoubleRegister(RegsHL))
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, false)

	return 8
}

func (c *Cpu) cpAIm() uint8 {
	valA := c.readRegister(RegA)
	im := c.currentByte()

	c.setFlag(FlagZ, valA-im == 0)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, doesByteSubHalfCarry(valA, im))
	c.setFlag(FlagC, doesByteSubCarry(valA, im))

	return 8
}

func (c *Cpu) andAIm() uint8 {
	valA := c.readRegister(RegA)
	im := c.currentByte()

	res := valA & im
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, true)
	c.setFlag(FlagC, false)

	return 8
}

func (c *Cpu) xorRegReg(a, b registerType) uint8 {
	res := c.readRegister(a) ^ c.readRegister(b)
	c.writeRegister(res, a)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, false)

	return 4
}

func (c *Cpu) xorADerefHL() uint8 {
	res := c.readRegister(RegA) ^ c.bus.ByteAt(c.readDoubleRegister(RegsHL))
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, false)

	return 8
}

func (c *Cpu) xorAIm() uint8 {
	valA := c.readRegister(RegA)
	im := c.currentByte()

	res := valA ^ im
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, false)

	return 8
}

func (c *Cpu) addAIm() uint8 {
	valA := c.readRegister(RegA)
	im := c.currentByte()

	res := valA + im
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, doesByteAddHalfCarry(valA, im))
	c.setFlag(FlagC, doesByteAddCarry(valA, im))

	return 8
}

func (c *Cpu) subAIm() uint8 {
	valA := c.readRegister(RegA)
	im := c.currentByte()

	res := valA - im
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, doesByteSubHalfCarry(valA, im))
	c.setFlag(FlagC, doesByteSubCarry(valA, im))

	return 8
}

func (c *Cpu) addADerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	b := c.bus.ByteAt(addr)
	a := c.readRegister(RegA)
	res := a + b
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, doesByteAddHalfCarry(a, b))
	c.setFlag(FlagC, doesByteAddCarry(a, b))

	return 8
}

func (c *Cpu) adcRegReg(a, b registerType) uint8 {
	aVal := c.readRegister(a)
	bVal := c.readRegister(b)

	var carry uint8
	if c.checkFlag(FlagC) {
		carry = 1
	}

	res := aVal + bVal + carry
	c.writeRegister(res, a)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)

	halfCarry := doesByteAddHalfCarry(aVal, bVal)
	if !halfCarry {
		halfCarry = doesByteAddHalfCarry(aVal+bVal, carry)
	}
	c.setFlag(FlagH, halfCarry)

	fullCarry := doesByteAddCarry(aVal, bVal)
	if !fullCarry {
		fullCarry = doesByteAddCarry(aVal+bVal, carry)
	}
	c.setFlag(FlagC, fullCarry)

	return 4
}

func (c *Cpu) adcADerefHL() uint8 {
	aVal := c.readRegister(RegA)
	addr := c.readDoubleRegister(RegsHL)
	bVal := c.bus.ByteAt(addr)

	var carry uint8
	if c.checkFlag(FlagC) {
		carry = 1
	}

	res := aVal + bVal + carry
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)

	halfCarry := doesByteAddHalfCarry(aVal, bVal)
	if !halfCarry {
		halfCarry = doesByteAddHalfCarry(aVal+bVal, carry)
	}
	c.setFlag(FlagH, halfCarry)

	fullCarry := doesByteAddCarry(aVal, bVal)
	if !fullCarry {
		fullCarry = doesByteAddCarry(aVal+bVal, carry)
	}
	c.setFlag(FlagC, fullCarry)

	return 8
}

func (c *Cpu) adcAIm() uint8 {
	aVal := c.readRegister(RegA)
	im := c.currentByte()

	var carry uint8
	if c.checkFlag(FlagC) {
		carry = 1
	}

	res := aVal + im + carry
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)

	halfCarry := doesByteAddHalfCarry(aVal, im)
	if !halfCarry {
		halfCarry = doesByteAddHalfCarry(aVal+im, carry)
	}
	c.setFlag(FlagH, halfCarry)

	fullCarry := doesByteAddCarry(aVal, im)
	if !fullCarry {
		fullCarry = doesByteAddCarry(aVal+im, carry)
	}
	c.setFlag(FlagC, fullCarry)

	return 8
}

func (c *Cpu) subRegReg(a, b registerType) uint8 {
	aVal := c.readRegister(a)
	bVal := c.readRegister(b)
	res := aVal - bVal
	c.writeRegister(res, a)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, doesByteSubHalfCarry(aVal, bVal))
	c.setFlag(FlagC, doesByteSubCarry(aVal, bVal))

	return 4
}

func (c *Cpu) subADerefHL() uint8 {
	aVal := c.readRegister(RegA)
	addr := c.readDoubleRegister(RegsHL)
	bVal := c.bus.ByteAt(addr)
	res := aVal - bVal
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, doesByteSubHalfCarry(aVal, bVal))
	c.setFlag(FlagC, doesByteSubCarry(aVal, bVal))

	return 8
}

func (c *Cpu) sbcRegReg(a, b registerType) uint8 {
	aVal := c.readRegister(a)
	bVal := c.readRegister(b)
	var carry uint8
	if c.checkFlag(FlagC) {
		carry = 1
	}
	res := aVal - bVal - carry
	c.writeRegister(res, a)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, true)

	halfCarry := doesByteSubHalfCarry(aVal, bVal)
	if !halfCarry {
		halfCarry = doesByteSubHalfCarry(aVal-bVal, carry)
	}
	c.setFlag(FlagH, halfCarry)

	fullCarry := doesByteSubCarry(aVal, bVal)
	if !fullCarry {
		fullCarry = doesByteSubCarry(aVal-bVal, carry)
	}
	c.setFlag(FlagC, fullCarry)

	return 4
}

func (c *Cpu) sbcADerefHL() uint8 {
	aVal := c.readRegister(RegA)
	addr := c.readDoubleRegister(RegsHL)
	bVal := c.bus.ByteAt(addr)
	var carry uint8
	if c.checkFlag(FlagC) {
		carry = 1
	}
	res := aVal - bVal - carry
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, true)

	halfCarry := doesByteSubHalfCarry(aVal, bVal)
	if !halfCarry {
		halfCarry = doesByteSubHalfCarry(aVal-bVal, carry)
	}
	c.setFlag(FlagH, halfCarry)

	fullCarry := doesByteSubCarry(aVal, bVal)
	if !fullCarry {
		fullCarry = doesByteSubCarry(aVal-bVal, carry)
	}
	c.setFlag(FlagC, fullCarry)

	return 8
}

func (c *Cpu) andRegReg(a, b registerType) uint8 {
	res := c.readRegister(a) & c.readRegister(b)
	c.writeRegister(res, a)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, true)
	c.setFlag(FlagC, false)

	return 4
}

func (c *Cpu) andADerefHL() uint8 {
	addr := c.readDoubleRegister(RegsHL)
	res := c.readRegister(RegA) & c.bus.ByteAt(addr)
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, true)
	c.setFlag(FlagC, false)

	return 8
}

func (c *Cpu) cpRegReg(a, b registerType) uint8 {
	valA := c.readRegister(a)
	valB := c.readRegister(b)

	c.setFlag(FlagZ, valA == valB)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, doesByteSubHalfCarry(valA, valB))
	c.setFlag(FlagC, doesByteSubCarry(valA, valB))

	return 4
}

func (c *Cpu) cpADerefHL() uint8 {
	valA := c.readRegister(RegA)
	addr := c.readDoubleRegister(RegsHL)
	valB := c.bus.ByteAt(addr)

	c.setFlag(FlagZ, valA == valB)
	c.setFlag(FlagN, true)
	c.setFlag(FlagH, doesByteSubHalfCarry(valA, valB))
	c.setFlag(FlagC, doesByteSubCarry(valA, valB))

	return 8
}

func (c *Cpu) addSPIm() uint8 {
	offset := int8(c.currentByte())
	oldSP := c.sp
	c.sp += word(offset)

	c.setFlag(FlagZ, false)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, (oldSP&0xf)+(word(offset)&0xf) > 0xf)
	c.setFlag(FlagC, (oldSP&0xff)+(word(offset)&0xff) > 0xff)

	return 16
}

func (c *Cpu) orAIm() uint8 {
	valA := c.readRegister(RegA)
	im := c.currentByte()

	res := valA | im
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, false)

	return 8
}

func (c *Cpu) sbcAIm() uint8 {
	aVal := c.readRegister(RegA)
	im := c.currentByte()

	var carry uint8
	if c.checkFlag(FlagC) {
		carry = 1
	}

	res := aVal - im - carry
	c.writeRegister(res, RegA)

	c.setFlag(FlagZ, res == 0)
	c.setFlag(FlagN, true)

	halfCarry := doesByteSubHalfCarry(aVal, im)
	if !halfCarry {
		halfCarry = doesByteSubHalfCarry(aVal-im, carry)
	}
	c.setFlag(FlagH, halfCarry)

	fullCarry := doesByteSubCarry(aVal, im)
	if !fullCarry {
		fullCarry = doesByteSubCarry(aVal-im, carry)
	}
	c.setFlag(FlagC, fullCarry)

	return 8
}
