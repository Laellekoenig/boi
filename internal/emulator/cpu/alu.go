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

type subAIm struct{}

func (i *subAIm) String() string { return "SUB A, u8" }
func (i *subAIm) Execute(c *Cpu) uint8 {
	return c.subAIm()
}

type addAB struct{}

func (i *addAB) String() string { return "ADD A, B" }
func (i *addAB) Execute(c *Cpu) uint8 {
	return c.addRegReg(RegA, RegB)
}

type addAC struct{}

func (i *addAC) String() string { return "ADD A, C" }
func (i *addAC) Execute(c *Cpu) uint8 {
	return c.addRegReg(RegA, RegC)
}

type addAD struct{}

func (i *addAD) String() string { return "ADD A, D" }
func (i *addAD) Execute(c *Cpu) uint8 {
	return c.addRegReg(RegA, RegD)
}

type addAE struct{}

func (i *addAE) String() string { return "ADD A, E" }
func (i *addAE) Execute(c *Cpu) uint8 {
	return c.addRegReg(RegA, RegE)
}

type addAH struct{}

func (i *addAH) String() string { return "ADD A, H" }
func (i *addAH) Execute(c *Cpu) uint8 {
	return c.addRegReg(RegA, RegH)
}

type addAL struct{}

func (i *addAL) String() string { return "ADD A, L" }
func (i *addAL) Execute(c *Cpu) uint8 {
	return c.addRegReg(RegA, RegL)
}

type addADerefHL struct{}

func (i *addADerefHL) String() string { return "ADD A, (HL)" }
func (i *addADerefHL) Execute(c *Cpu) uint8 {
	return c.addADerefHL()
}

type addAA struct{}

func (i *addAA) String() string { return "ADD A, A" }
func (i *addAA) Execute(c *Cpu) uint8 {
	return c.addRegReg(RegA, RegA)
}

type adcAB struct{}

func (i *adcAB) String() string { return "ADC A, B" }
func (i *adcAB) Execute(c *Cpu) uint8 {
	return c.adcRegReg(RegA, RegB)
}

type adcAC struct{}

func (i *adcAC) String() string { return "ADC A, C" }
func (i *adcAC) Execute(c *Cpu) uint8 {
	return c.adcRegReg(RegA, RegC)
}

type adcAD struct{}

func (i *adcAD) String() string { return "ADC A, D" }
func (i *adcAD) Execute(c *Cpu) uint8 {
	return c.adcRegReg(RegA, RegD)
}

type adcAE struct{}

func (i *adcAE) String() string { return "ADC A, E" }
func (i *adcAE) Execute(c *Cpu) uint8 {
	return c.adcRegReg(RegA, RegE)
}

type adcAH struct{}

func (i *adcAH) String() string { return "ADC A, H" }
func (i *adcAH) Execute(c *Cpu) uint8 {
	return c.adcRegReg(RegA, RegH)
}

type adcAL struct{}

func (i *adcAL) String() string { return "ADC A, L" }
func (i *adcAL) Execute(c *Cpu) uint8 {
	return c.adcRegReg(RegA, RegL)
}

type adcADerefHL struct{}

func (i *adcADerefHL) String() string { return "ADC A, (HL)" }
func (i *adcADerefHL) Execute(c *Cpu) uint8 {
	return c.adcADerefHL()
}

type adcAA struct{}

func (i *adcAA) String() string { return "ADC A, A" }
func (i *adcAA) Execute(c *Cpu) uint8 {
	return c.adcRegReg(RegA, RegA)
}

type subAB struct{}

func (i *subAB) String() string { return "SUB A, B" }
func (i *subAB) Execute(c *Cpu) uint8 {
	return c.subRegReg(RegA, RegB)
}

type subAC struct{}

func (i *subAC) String() string { return "SUB A, C" }
func (i *subAC) Execute(c *Cpu) uint8 {
	return c.subRegReg(RegA, RegC)
}

type subAD struct{}

func (i *subAD) String() string { return "SUB A, D" }
func (i *subAD) Execute(c *Cpu) uint8 {
	return c.subRegReg(RegA, RegD)
}

type subAE struct{}

func (i *subAE) String() string { return "SUB A, E" }
func (i *subAE) Execute(c *Cpu) uint8 {
	return c.subRegReg(RegA, RegE)
}

type subAH struct{}

func (i *subAH) String() string { return "SUB A, H" }
func (i *subAH) Execute(c *Cpu) uint8 {
	return c.subRegReg(RegA, RegH)
}

type subAL struct{}

func (i *subAL) String() string { return "SUB A, L" }
func (i *subAL) Execute(c *Cpu) uint8 {
	return c.subRegReg(RegA, RegL)
}

type subADerefHL struct{}

func (i *subADerefHL) String() string { return "SUB A, (HL)" }
func (i *subADerefHL) Execute(c *Cpu) uint8 {
	return c.subADerefHL()
}

type subAA struct{}

func (i *subAA) String() string { return "SUB A, A" }
func (i *subAA) Execute(c *Cpu) uint8 {
	return c.subRegReg(RegA, RegA)
}

type sbcAB struct{}

func (i *sbcAB) String() string { return "SBC A, B" }
func (i *sbcAB) Execute(c *Cpu) uint8 {
	return c.sbcRegReg(RegA, RegB)
}

type sbcAC struct{}

func (i *sbcAC) String() string { return "SBC A, C" }
func (i *sbcAC) Execute(c *Cpu) uint8 {
	return c.sbcRegReg(RegA, RegC)
}

type sbcAD struct{}

func (i *sbcAD) String() string { return "SBC A, D" }
func (i *sbcAD) Execute(c *Cpu) uint8 {
	return c.sbcRegReg(RegA, RegD)
}

type sbcAE struct{}

func (i *sbcAE) String() string { return "SBC A, E" }
func (i *sbcAE) Execute(c *Cpu) uint8 {
	return c.sbcRegReg(RegA, RegE)
}

type sbcAH struct{}

func (i *sbcAH) String() string { return "SBC A, H" }
func (i *sbcAH) Execute(c *Cpu) uint8 {
	return c.sbcRegReg(RegA, RegH)
}

type sbcAL struct{}

func (i *sbcAL) String() string { return "SBC A, L" }
func (i *sbcAL) Execute(c *Cpu) uint8 {
	return c.sbcRegReg(RegA, RegL)
}

type sbcADerefHL struct{}

func (i *sbcADerefHL) String() string { return "SBC A, (HL)" }
func (i *sbcADerefHL) Execute(c *Cpu) uint8 {
	return c.sbcADerefHL()
}

type sbcAA struct{}

func (i *sbcAA) String() string { return "SBC A, A" }
func (i *sbcAA) Execute(c *Cpu) uint8 {
	return c.sbcRegReg(RegA, RegA)
}

type andAB struct{}

func (i *andAB) String() string { return "AND A, B" }
func (i *andAB) Execute(c *Cpu) uint8 {
	return c.andRegReg(RegA, RegB)
}

type andAC struct{}

func (i *andAC) String() string { return "AND A, C" }
func (i *andAC) Execute(c *Cpu) uint8 {
	return c.andRegReg(RegA, RegC)
}

type andAD struct{}

func (i *andAD) String() string { return "AND A, D" }
func (i *andAD) Execute(c *Cpu) uint8 {
	return c.andRegReg(RegA, RegD)
}

type andAE struct{}

func (i *andAE) String() string { return "AND A, E" }
func (i *andAE) Execute(c *Cpu) uint8 {
	return c.andRegReg(RegA, RegE)
}

type andAH struct{}

func (i *andAH) String() string { return "AND A, H" }
func (i *andAH) Execute(c *Cpu) uint8 {
	return c.andRegReg(RegA, RegH)
}

type andAL struct{}

func (i *andAL) String() string { return "AND A, L" }
func (i *andAL) Execute(c *Cpu) uint8 {
	return c.andRegReg(RegA, RegL)
}

type andADerefHL struct{}

func (i *andADerefHL) String() string { return "AND A, (HL)" }
func (i *andADerefHL) Execute(c *Cpu) uint8 {
	return c.andADerefHL()
}

type andAA struct{}

func (i *andAA) String() string { return "AND A, A" }
func (i *andAA) Execute(c *Cpu) uint8 {
	return c.andRegReg(RegA, RegA)
}

type cpAB struct{}

func (i *cpAB) String() string { return "CP A, B" }
func (i *cpAB) Execute(c *Cpu) uint8 {
	return c.cpRegReg(RegA, RegB)
}

type cpAC struct{}

func (i *cpAC) String() string { return "CP A, C" }
func (i *cpAC) Execute(c *Cpu) uint8 {
	return c.cpRegReg(RegA, RegC)
}

type cpAD struct{}

func (i *cpAD) String() string { return "CP A, D" }
func (i *cpAD) Execute(c *Cpu) uint8 {
	return c.cpRegReg(RegA, RegD)
}

type cpAE struct{}

func (i *cpAE) String() string { return "CP A, E" }
func (i *cpAE) Execute(c *Cpu) uint8 {
	return c.cpRegReg(RegA, RegE)
}

type cpAH struct{}

func (i *cpAH) String() string { return "CP A, H" }
func (i *cpAH) Execute(c *Cpu) uint8 {
	return c.cpRegReg(RegA, RegH)
}

type cpAL struct{}

func (i *cpAL) String() string { return "CP A, L" }
func (i *cpAL) Execute(c *Cpu) uint8 {
	return c.cpRegReg(RegA, RegL)
}

type cpADerefHL struct{}

func (i *cpADerefHL) String() string { return "CP A, (HL)" }
func (i *cpADerefHL) Execute(c *Cpu) uint8 {
	return c.cpADerefHL()
}

type cpAA struct{}

func (i *cpAA) String() string { return "CP A, A" }
func (i *cpAA) Execute(c *Cpu) uint8 {
	return c.cpRegReg(RegA, RegA)
}

type addSPIm struct{}

func (i *addSPIm) String() string { return "ADD SP, i8" }
func (i *addSPIm) Execute(c *Cpu) uint8 {
	return c.addSPIm()
}

type orAIm struct{}

func (i *orAIm) String() string { return "OR A, u8" }
func (i *orAIm) Execute(c *Cpu) uint8 {
	return c.orAIm()
}

type adcAIm struct{}

func (i *adcAIm) String() string { return "ADC A, u8" }
func (i *adcAIm) Execute(c *Cpu) uint8 {
	return c.adcAIm()
}
