package cpu

func (c *Cpu) ldRegsIm(t doubleRegisterType) {
	w := c.currentWord()
	c.writeDoubleRegister(w, t)
}

func (c *Cpu) ldDerefImSP() {
	c.bus.WriteWordAt(c.sp, c.currentWord())
}

func (c *Cpu) ldRegsRegs(from, to doubleRegisterType) {
	c.writeDoubleRegister(c.readDoubleRegister(from), to)
}

func (c *Cpu) ldDerefRegsReg(from registerType, to doubleRegisterType) {
	addr := c.readDoubleRegister(to)
	val := c.readRegister(from)
	c.bus.WriteByteAt(val, addr)
}

func (c *Cpu) ldRegIm(t registerType) {
	c.writeRegister(c.currentByte(), t)
}

func (c *Cpu) ldRegDerefRegs(from doubleRegisterType, to registerType) {
	addr := c.readDoubleRegister(from)
	c.writeRegister(c.bus.ByteAt(addr), to)
}

func (c *Cpu) ldDerefIncHLA() {
	addr := c.readDoubleRegister(RegsHL)
	c.incRegs(RegsHL)
	c.bus.WriteByteAt(c.readRegister(RegA), addr)
}

func (c *Cpu) ldDerefDecHLA() {
	addr := c.readDoubleRegister(RegsHL)
	c.decRegs(RegsHL)
	c.bus.WriteByteAt(c.readRegister(RegA), addr)
}

func (c *Cpu) ldADerefIncHL() {
	addr := c.readDoubleRegister(RegsHL)
	c.incRegs(RegsHL)
	c.writeRegister(c.bus.ByteAt(addr), RegA)
}

func (c *Cpu) ldADerefDecHL() {
	addr := c.readDoubleRegister(RegsHL)
	c.decRegs(RegsHL)
	c.writeRegister(c.bus.ByteAt(addr), RegA)
}

func (c *Cpu) ldRegReg(from, to registerType) {
	c.writeRegister(c.readRegister(from), to)
}

func (c *Cpu) ldOffsetImA() {
	addr := 0xff00 + word(c.currentByte())
	c.bus.WriteByteAt(c.readRegister(RegA), addr)
}

func (c *Cpu) ldOffsetCA() {
	addr := 0xff00 + word(c.readRegister(RegC))
	c.bus.WriteByteAt(c.readRegister(RegA), addr)
}

func (c *Cpu) ldAOffsetIm() {
	addr := 0xff00 + word(c.currentByte())
	c.writeRegister(c.bus.ByteAt(addr), RegA)
}

func (c *Cpu) ldAOffsetC() {
	addr := 0xff00 + word(c.readRegister(RegC))
	c.writeRegister(c.bus.ByteAt(addr), RegA)
}

func (c *Cpu) ldDerefImA() {
	c.bus.WriteWordAt(c.currentWord(), c.currentWord())
}

func (c *Cpu) ldADerefIm() {
	addr := c.currentWord()
	c.writeRegister(c.bus.ByteAt(addr), RegA)
}

func (c *Cpu) ldDerefHLIm() {
	addr := c.readDoubleRegister(RegsHL)
	c.bus.WriteByteAt(c.currentByte(), addr)
}

func (c *Cpu) ldSPHL() {
	c.sp = c.readDoubleRegister(RegsHL)
}

func (c *Cpu) ldSPIm() {
	c.sp = c.currentWord()
}
