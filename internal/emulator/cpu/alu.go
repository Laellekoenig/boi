package cpu

type incBC struct{}

func (i *incBC) String() string { return "INC BC" }
func (i *incBC) Execute(c *Cpu) uint8 {
	c.incRegs(RegsBC)
	return 8
}

type incB struct{}

func (i *incB) String() string { return "INC B" }
func (i *incB) Execute(c *Cpu) uint8 {
	c.incReg(RegB)
	return 4
}

type decB struct{}

func (i *decB) String() string { return "DEC B" }
func (i *decB) Execute(c *Cpu) uint8 {
	c.decReg(RegB)
	return 4
}

type addHLBC struct{}

func (i *addHLBC) String() string { return "ADD HL, BC" }
func (i *addHLBC) Execute(c *Cpu) uint8 {
	c.addRegsRegs(RegsHL, RegsBC)
	return 8
}

type decBC struct{}

func (i *decBC) String() string { return "DEC BC" }
func (i *decBC) Execute(c *Cpu) uint8 {
	c.decRegs(RegsBC)
	return 8
}

type incC struct{}

func (i *incC) String() string { return "INC C" }
func (i *incC) Execute(c *Cpu) uint8 {
	c.incReg(RegC)
	return 4
}

type decC struct{}

func (i *decC) String() string { return "DEC C" }
func (i *decC) Execute(c *Cpu) uint8 {
	c.decReg(RegC)
	return 4
}

type incDE struct{}

func (i *incDE) String() string { return "INC DE" }
func (i *incDE) Execute(c *Cpu) uint8 {
	c.incRegs(RegsDE)
	return 8
}

type incD struct{}

func (i *incD) String() string { return "INC D" }
func (i *incD) Execute(c *Cpu) uint8 {
	c.incReg(RegD)
	return 4
}

type decD struct{}

func (i *decD) String() string { return "DEC D" }
func (i *decD) Execute(c *Cpu) uint8 {
	c.decReg(RegD)
	return 4
}

type addHLDE struct{}

func (i *addHLDE) String() string { return "ADD HL, DE" }
func (i *addHLDE) Execute(c *Cpu) uint8 {
	c.addRegsRegs(RegsHL, RegsDE)
	return 8
}

type decDE struct{}

func (i *decDE) String() string { return "DEC DE" }
func (i *decDE) Execute(c *Cpu) uint8 {
	c.decRegs(RegsDE)
	return 8
}

type incE struct{}

func (i *incE) String() string { return "INC E" }
func (i *incE) Execute(c *Cpu) uint8 {
	c.incReg(RegE)
	return 4
}

type decE struct{}

func (i *decE) String() string { return "DEC E" }
func (i *decE) Execute(c *Cpu) uint8 {
	c.decReg(RegE)
	return 4
}

type incHL struct{}

func (i *incHL) String() string { return "INC HL" }
func (i *incHL) Execute(c *Cpu) uint8 {
	c.incRegs(RegsHL)
	return 8
}

type incH struct{}

func (i *incH) String() string { return "INC H" }
func (i *incH) Execute(c *Cpu) uint8 {
	c.incReg(RegH)
	return 4
}

type decH struct{}

func (i *decH) String() string { return "DEC H" }
func (i *decH) Execute(c *Cpu) uint8 {
	c.decReg(RegH)
	return 4
}

type daa struct{}

func (i *daa) String() string { return "DAA" }
func (i *daa) Execute(c *Cpu) uint8 {
	c.daa()
	return 4
}

type addHLHL struct{}

func (i *addHLHL) String() string { return "ADD HL, HL" }
func (i *addHLHL) Execute(c *Cpu) uint8 {
	c.addRegsRegs(RegsHL, RegsHL)
	return 8
}

type decHL struct{}

func (i *decHL) String() string { return "DEC HL" }
func (i *decHL) Execute(c *Cpu) uint8 {
	c.decRegs(RegsHL)
	return 8
}

type incL struct{}

func (i *incL) String() string { return "INC L" }
func (i *incL) Execute(c *Cpu) uint8 {
	c.incReg(RegL)
	return 4
}

type decL struct{}

func (i *decL) String() string { return "DEC L" }
func (i *decL) Execute(c *Cpu) uint8 {
	c.decReg(RegL)
	return 4
}

type cpl struct{}

func (i *cpl) String() string { return "CPL" }
func (i *cpl) Execute(c *Cpu) uint8 {
	c.cpl()
	return 4
}

type incSP struct{}

func (i *incSP) String() string { return "INC SP" }
func (i *incSP) Execute(c *Cpu) uint8 {
	c.sp += 1
	return 8
}

type incDerefHL struct{}

func (i *incDerefHL) String() string { return "INC (HL)" }
func (i *incDerefHL) Execute(c *Cpu) uint8 {
	c.incDerefHL()
	return 12
}

type decDerefHL struct{}

func (i *decDerefHL) String() string { return "DEC (HL)" }
func (i *decDerefHL) Execute(c *Cpu) uint8 {
	c.decDerefHL()
	return 12
}

type scf struct{}

func (i *scf) String() string { return "SCF" }
func (i *scf) Execute(c *Cpu) uint8 {
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, true)
	return 4
}

type addHLSP struct{}

func (i *addHLSP) String() string { return "ADD HL, SP" }
func (i *addHLSP) Execute(c *Cpu) uint8 {
	c.addHLSP()
	return 8
}

type decSP struct{}

func (i *decSP) String() string { return "DEC SP" }
func (i *decSP) Execute(c *Cpu) uint8 {
	c.sp -= 1
	return 8
}

type incA struct{}

func (i *incA) String() string { return "INC A" }
func (i *incA) Execute(c *Cpu) uint8 {
	c.incReg(RegA)
	return 4
}

type decA struct{}

func (i *decA) String() string { return "DEC A" }
func (i *decA) Execute(c *Cpu) uint8 {
	c.decReg(RegA)
	return 4
}

type ccf struct{}

func (i *ccf) String() string { return "CCF" }
func (i *ccf) Execute(c *Cpu) uint8 {
	c.setFlag(FlagN, false)
	c.setFlag(FlagH, false)
	c.setFlag(FlagC, !c.checkFlag(FlagC))
	return 4
}

type orAB struct{}

func (i *orAB) String() string { return "OR A, B" }
func (i *orAB) Execute(c *Cpu) uint8 {
	return c.orRegReg(RegA, RegB)
}

type orAC struct{}

func (i *orAC) String() string { return "OR A, C" }
func (i *orAC) Execute(c *Cpu) uint8 {
	return c.orRegReg(RegA, RegC)
}

type orAD struct{}

func (i *orAD) String() string { return "OR A, D" }
func (i *orAD) Execute(c *Cpu) uint8 {
	return c.orRegReg(RegA, RegD)
}

type orAE struct{}

func (i *orAE) String() string { return "OR A, E" }
func (i *orAE) Execute(c *Cpu) uint8 {
	return c.orRegReg(RegA, RegE)
}

type orAH struct{}

func (i *orAH) String() string { return "OR A, H" }
func (i *orAH) Execute(c *Cpu) uint8 {
	return c.orRegReg(RegA, RegH)
}

type orAL struct{}

func (i *orAL) String() string { return "OR A, L" }
func (i *orAL) Execute(c *Cpu) uint8 {
	return c.orRegReg(RegA, RegL)
}

type orADerefHL struct{}

func (i *orADerefHL) String() string { return "OR A, (HL)" }
func (i *orADerefHL) Execute(c *Cpu) uint8 {
	return c.orADerefHL()
}

type orAA struct{}

func (i *orAA) String() string { return "OR A, A" }
func (i *orAA) Execute(c *Cpu) uint8 {
	return c.orRegReg(RegA, RegA)
}

type cpAIm struct{}

func (i *cpAIm) String() string { return "CP A, u8" }
func (i *cpAIm) Execute(c *Cpu) uint8 {
	return c.cpAIm()
}

type andAIm struct{}

func (i *andAIm) String() string { return "AND A, u8" }
func (i *andAIm) Execute(c *Cpu) uint8 {
	return c.andAIm()
}

type xorAB struct{}

func (i *xorAB) String() string { return "XOR A, B" }
func (i *xorAB) Execute(c *Cpu) uint8 {
	return c.xorRegReg(RegA, RegB)
}

type xorAC struct{}

func (i *xorAC) String() string { return "XOR A, C" }
func (i *xorAC) Execute(c *Cpu) uint8 {
	return c.xorRegReg(RegA, RegC)
}

type xorAD struct{}

func (i *xorAD) String() string { return "XOR A, D" }
func (i *xorAD) Execute(c *Cpu) uint8 {
	return c.xorRegReg(RegA, RegD)
}

type xorAE struct{}

func (i *xorAE) String() string { return "XOR A, E" }
func (i *xorAE) Execute(c *Cpu) uint8 {
	return c.xorRegReg(RegA, RegE)
}

type xorAH struct{}

func (i *xorAH) String() string { return "XOR A, H" }
func (i *xorAH) Execute(c *Cpu) uint8 {
	return c.xorRegReg(RegA, RegH)
}

type xorAL struct{}

func (i *xorAL) String() string { return "XOR A, L" }
func (i *xorAL) Execute(c *Cpu) uint8 {
	return c.xorRegReg(RegA, RegL)
}

type xorADerefHL struct{}

func (i *xorADerefHL) String() string { return "XOR A, (HL)" }
func (i *xorADerefHL) Execute(c *Cpu) uint8 {
	return c.xorADerefHL()
}

type xorAA struct{}

func (i *xorAA) String() string { return "XOR A, A" }
func (i *xorAA) Execute(c *Cpu) uint8 {
	return c.xorRegReg(RegA, RegA)
}

type xorAIm struct{}

func (i *xorAIm) String() string { return "XOR A, u8" }
func (i *xorAIm) Execute(c *Cpu) uint8 {
	return c.xorAIm()
}

type addAIm struct{}

func (i *addAIm) String() string { return "ADD A, u8" }
func (i *addAIm) Execute(c *Cpu) uint8 {
	return c.addAIm()
}
