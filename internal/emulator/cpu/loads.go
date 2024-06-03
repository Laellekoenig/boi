package cpu

type ldBCIm struct{}

func (l *ldBCIm) String() string { return "LD BC, u16" }
func (l *ldBCIm) Execute(c *Cpu) uint8 {
	c.ldRegsIm(RegsBC)
	return 12
}

type ldDerefBCA struct{}

func (l *ldDerefBCA) String() string { return "LD (BC), A" }
func (l *ldDerefBCA) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegA, RegsBC)
	return 8
}

type ldBIm struct{}

func (l *ldBIm) String() string { return "LD B, u8" }
func (l *ldBIm) Execute(c *Cpu) uint8 {
	c.ldRegIm(RegB)
	return 8
}

type ldDerefImSP struct{}

func (l *ldDerefImSP) String() string { return "LD (u16), SP" }
func (l *ldDerefImSP) Execute(c *Cpu) uint8 {
	c.ldDerefImSP()
	return 20
}

type ldADerefBC struct{}

func (l *ldADerefBC) String() string { return "LD A, (BC)" }
func (l *ldADerefBC) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsBC, RegA)
	return 8
}

type ldCIm struct{}

func (l *ldCIm) String() string { return "LD C, u8" }
func (l *ldCIm) Execute(c *Cpu) uint8 {
	c.ldRegIm(RegC)
	return 8
}

type ldDEIm struct{}

func (l *ldDEIm) String() string { return "LD DE, u16" }
func (l *ldDEIm) Execute(c *Cpu) uint8 {
	c.ldRegsIm(RegsDE)
	return 12
}

type ldDerefDEA struct{}

func (l *ldDerefDEA) String() string { return "LD (DE), A" }
func (l *ldDerefDEA) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegA, RegsDE)
	return 8
}

type ldDIm struct{}

func (l *ldDIm) String() string { return "LD D, u8" }
func (l *ldDIm) Execute(c *Cpu) uint8 {
	c.ldRegIm(RegD)
	return 8
}

type ldADerefDE struct{}

func (l *ldADerefDE) String() string { return "LD A, (DE)" }
func (l *ldADerefDE) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsDE, RegA)
	return 8
}

type ldEIm struct{}

func (l *ldEIm) String() string { return "LD E, u8" }
func (l *ldEIm) Execute(c *Cpu) uint8 {
	c.ldRegIm(RegE)
	return 8
}

type ldHLIm struct{}

func (l *ldHLIm) String() string { return "LD HL, u16" }
func (l *ldHLIm) Execute(c *Cpu) uint8 {
	c.ldRegsIm(RegsHL)
	return 12
}

type ldDerefIncHLA struct{}

func (l *ldDerefIncHLA) String() string { return "LD (HL+), A" }
func (l *ldDerefIncHLA) Execute(c *Cpu) uint8 {
	c.ldDerefIncHLA()
	return 8
}

type ldHIm struct{}

func (l *ldHIm) String() string { return "LD H, u8" }
func (l *ldHIm) Execute(c *Cpu) uint8 {
	c.ldRegIm(RegH)
	return 8
}

type ldADerefIncHL struct{}

func (l *ldADerefIncHL) String() string { return "LD A, (HL+)" }
func (l *ldADerefIncHL) Execute(c *Cpu) uint8 {
	c.ldADerefIncHL()
	return 8
}

type ldLIm struct{}

func (l *ldLIm) String() string { return "LD L, u8" }
func (l *ldLIm) Execute(c *Cpu) uint8 {
	c.ldRegIm(RegL)
	return 8
}

type ldSPIm struct{}

func (l *ldSPIm) String() string { return "LD SP, u16" }
func (l *ldSPIm) Execute(c *Cpu) uint8 {
	c.ldSPIm()
	return 12
}

type ldDerefDecHLA struct{}

func (l *ldDerefDecHLA) String() string { return "LD (HL-), A" }
func (l *ldDerefDecHLA) Execute(c *Cpu) uint8 {
	c.ldDerefDecHLA()
	return 8
}

type ldDerefHLIm struct{}

func (l *ldDerefHLIm) String() string { return "LD (HL), u8" }
func (l *ldDerefHLIm) Execute(c *Cpu) uint8 {
	c.ldDerefHLIm()
	return 12
}

type ldADerefDecHL struct{}

func (l *ldADerefDecHL) String() string { return "LD A, (HL-)" }
func (l *ldADerefDecHL) Execute(c *Cpu) uint8 {
	c.ldADerefDecHL()
	return 8
}

type ldAIm struct{}

func (l *ldAIm) String() string { return "LD A, u8" }
func (l *ldAIm) Execute(c *Cpu) uint8 {
	c.ldRegIm(RegA)
	return 8
}

type ldBB struct{}

func (l *ldBB) String() string { return "LD B, B" }
func (l *ldBB) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegB, RegB)
	return 4
}

type ldBC struct{}

func (l *ldBC) String() string { return "LD B, C" }
func (l *ldBC) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegC, RegB)
	return 4
}

type ldBD struct{}

func (l *ldBD) String() string { return "LD B, D" }
func (l *ldBD) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegD, RegB)
	return 4
}

type ldBE struct{}

func (l *ldBE) String() string { return "LD B, E" }
func (l *ldBE) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegE, RegB)
	return 4
}

type ldBH struct{}

func (l *ldBH) String() string { return "LD B, H" }
func (l *ldBH) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegH, RegB)
	return 4
}

type ldBL struct{}

func (l *ldBL) String() string { return "LD B, L" }
func (l *ldBL) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegL, RegB)
	return 4
}

type ldBDerefHL struct{}

func (l *ldBDerefHL) String() string { return "LD B, (HL)" }
func (l *ldBDerefHL) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsHL, RegB)
	return 8
}

type ldBA struct{}

func (l *ldBA) String() string { return "LD B, A" }
func (l *ldBA) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegA, RegB)
	return 4
}

type ldCB struct{}

func (l *ldCB) String() string { return "LD C, B" }
func (l *ldCB) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegB, RegC)
	return 4
}

type ldCC struct{}

func (l *ldCC) String() string { return "LD C, C" }
func (l *ldCC) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegC, RegC)
	return 4
}

type ldCD struct{}

func (l *ldCD) String() string { return "LD C, D" }
func (l *ldCD) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegD, RegC)
	return 4
}

type ldCE struct{}

func (l *ldCE) String() string { return "LD C, E" }
func (l *ldCE) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegE, RegC)
	return 4
}

type ldCH struct{}

func (l *ldCH) String() string { return "LD C, H" }
func (l *ldCH) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegH, RegC)
	return 4
}

type ldCL struct{}

func (l *ldCL) String() string { return "LD C, L" }
func (l *ldCL) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegL, RegC)
	return 4
}

type ldCDerefHL struct{}

func (l *ldCDerefHL) String() string { return "LD C, (HL)" }
func (l *ldCDerefHL) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsHL, RegC)
	return 8
}

type ldCA struct{}

func (l *ldCA) String() string { return "LD C, A" }
func (l *ldCA) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegA, RegC)
	return 4
}

type ldDB struct{}

func (l *ldDB) String() string { return "LD D, B" }
func (l *ldDB) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegB, RegD)
	return 4
}

type ldDC struct{}

func (l *ldDC) String() string { return "LD D, C" }
func (l *ldDC) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegC, RegD)
	return 4
}

type ldDD struct{}

func (l *ldDD) String() string { return "LD D, D" }
func (l *ldDD) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegD, RegD)
	return 4
}

type ldDE struct{}

func (l *ldDE) String() string { return "LD D, E" }
func (l *ldDE) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegE, RegD)
	return 4
}

type ldDH struct{}

func (l *ldDH) String() string { return "LD D, H" }
func (l *ldDH) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegH, RegD)
	return 4
}

type ldDL struct{}

func (l *ldDL) String() string { return "LD D, L" }
func (l *ldDL) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegL, RegD)
	return 4
}

type ldDDerefHL struct{}

func (l *ldDDerefHL) String() string { return "LD D, (HL)" }
func (l *ldDDerefHL) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsHL, RegD)
	return 8
}

type ldDA struct{}

func (l *ldDA) String() string { return "LD D, A" }
func (l *ldDA) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegA, RegD)
	return 4
}

type ldEB struct{}

func (l *ldEB) String() string { return "LD E, B" }
func (l *ldEB) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegB, RegE)
	return 4
}

type ldEC struct{}

func (l *ldEC) String() string { return "LD E, C" }
func (l *ldEC) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegC, RegE)
	return 4
}

type ldED struct{}

func (l *ldED) String() string { return "LD E, D" }
func (l *ldED) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegD, RegE)
	return 4
}

type ldEE struct{}

func (l *ldEE) String() string { return "LD E, E" }
func (l *ldEE) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegE, RegE)
	return 4
}

type ldEH struct{}

func (l *ldEH) String() string { return "LD E, H" }
func (l *ldEH) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegH, RegE)
	return 4
}

type ldEL struct{}

func (l *ldEL) String() string { return "LD E, L" }
func (l *ldEL) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegL, RegE)
	return 4
}

type ldEDerefHL struct{}

func (l *ldEDerefHL) String() string { return "LD E, (HL)" }
func (l *ldEDerefHL) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsHL, RegE)
	return 8
}

type ldEA struct{}

func (l *ldEA) String() string { return "LD E, A" }
func (l *ldEA) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegA, RegE)
	return 4
}

type ldHB struct{}

func (l *ldHB) String() string { return "LD H, B" }
func (l *ldHB) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegB, RegH)
	return 4
}

type ldHC struct{}

func (l *ldHC) String() string { return "LD H, C" }
func (l *ldHC) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegC, RegH)
	return 4
}

type ldHD struct{}

func (l *ldHD) String() string { return "LD H, D" }
func (l *ldHD) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegD, RegH)
	return 4
}

type ldHE struct{}

func (l *ldHE) String() string { return "LD H, E" }
func (l *ldHE) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegE, RegH)
	return 4
}

type ldHH struct{}

func (l *ldHH) String() string { return "LD H, H" }
func (l *ldHH) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegH, RegH)
	return 4
}

type ldHL struct{}

func (l *ldHL) String() string { return "LD H, L" }
func (l *ldHL) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegL, RegH)
	return 4
}

type ldHDerefHL struct{}

func (l *ldHDerefHL) String() string { return "LD H, (HL)" }
func (l *ldHDerefHL) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsHL, RegH)
	return 8
}

type ldHA struct{}

func (l *ldHA) String() string { return "LD H, A" }
func (l *ldHA) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegA, RegH)
	return 4
}

type ldLB struct{}

func (l *ldLB) String() string { return "LD L, B" }
func (l *ldLB) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegB, RegL)
	return 4
}

type ldLC struct{}

func (l *ldLC) String() string { return "LD L, C" }
func (l *ldLC) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegC, RegL)
	return 4
}

type ldLD struct{}

func (l *ldLD) String() string { return "LD L, D" }
func (l *ldLD) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegD, RegL)
	return 4
}

type ldLE struct{}

func (l *ldLE) String() string { return "LD L, E" }
func (l *ldLE) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegE, RegL)
	return 4
}

type ldLH struct{}

func (l *ldLH) String() string { return "LD L, H" }
func (l *ldLH) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegH, RegL)
	return 4
}

type ldLL struct{}

func (l *ldLL) String() string { return "LD L, L" }
func (l *ldLL) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegL, RegL)
	return 4
}

type ldLDerefHL struct{}

func (l *ldLDerefHL) String() string { return "LD L, (HL)" }
func (l *ldLDerefHL) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsHL, RegL)
	return 8
}

type ldLA struct{}

func (l *ldLA) String() string { return "LD L, A" }
func (l *ldLA) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegA, RegL)
	return 4
}

type ldDerefHLB struct{}

func (l *ldDerefHLB) String() string { return "LD (HL), B" }
func (l *ldDerefHLB) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegB, RegsHL)
	return 8
}

type ldDerefHLC struct{}

func (l *ldDerefHLC) String() string { return "LD (HL), C" }
func (l *ldDerefHLC) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegC, RegsHL)
	return 8
}

type ldDerefHLD struct{}

func (l *ldDerefHLD) String() string { return "LD (HL), D" }
func (l *ldDerefHLD) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegD, RegsHL)
	return 8
}

type ldDerefHLE struct{}

func (l *ldDerefHLE) String() string { return "LD (HL), E" }
func (l *ldDerefHLE) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegE, RegsHL)
	return 8
}

type ldDerefHLH struct{}

func (l *ldDerefHLH) String() string { return "LD (HL), H" }
func (l *ldDerefHLH) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegH, RegsHL)
	return 8
}

type ldDerefHLL struct{}

func (l *ldDerefHLL) String() string { return "LD (HL), L" }
func (l *ldDerefHLL) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegL, RegsHL)
	return 8
}

type ldDerefHLA struct{}

func (l *ldDerefHLA) String() string { return "LD (HL), A" }
func (l *ldDerefHLA) Execute(c *Cpu) uint8 {
	c.ldDerefRegsReg(RegA, RegsHL)
	return 8
}

type ldAB struct{}

func (l *ldAB) String() string { return "LD A, B" }
func (l *ldAB) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegB, RegA)
	return 4
}

type ldAC struct{}

func (l *ldAC) String() string { return "LD A, C" }
func (l *ldAC) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegC, RegA)
	return 4
}

type ldAD struct{}

func (l *ldAD) String() string { return "LD A, D" }
func (l *ldAD) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegD, RegA)
	return 4
}

type ldAE struct{}

func (l *ldAE) String() string { return "LD A, E" }
func (l *ldAE) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegE, RegA)
	return 4
}

type ldAH struct{}

func (l *ldAH) String() string { return "LD A, H" }
func (l *ldAH) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegH, RegA)
	return 4
}

type ldAL struct{}

func (l *ldAL) String() string { return "LD A, L" }
func (l *ldAL) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegL, RegA)
	return 4
}

type ldADerefHL struct{}

func (l *ldADerefHL) String() string { return "LD A, (HL)" }
func (l *ldADerefHL) Execute(c *Cpu) uint8 {
	c.ldRegDerefRegs(RegsHL, RegA)
	return 8
}

type ldAA struct{}

func (l *ldAA) String() string { return "LD A, A" }
func (l *ldAA) Execute(c *Cpu) uint8 {
	c.ldRegReg(RegA, RegA)
	return 4
}

type ldDerefOffsetImA struct{}

func (l *ldDerefOffsetImA) String() string { return "LD (0xff00 + u8), A" }
func (l *ldDerefOffsetImA) Execute(c *Cpu) uint8 {
	c.ldOffsetImA()
	return 12
}

type ldDerefOffsetCA struct{}

func (l *ldDerefOffsetCA) String() string { return "LD (0xff00 + C), A" }
func (l *ldDerefOffsetCA) Execute(c *Cpu) uint8 {
	c.ldOffsetCA()
	return 8
}

type ldDerefImA struct{}

func (l *ldDerefImA) String() string { return "LD (u16), A" }
func (l *ldDerefImA) Execute(c *Cpu) uint8 {
	c.ldDerefImA()
	return 16
}

type ldADerefOffsetIm struct{}

func (l *ldADerefOffsetIm) String() string { return "LD A, (0xff00 + u8)" }
func (l *ldADerefOffsetIm) Execute(c *Cpu) uint8 {
	c.ldAOffsetIm()
	return 12
}

type ldADerefOffsetC struct{}

func (l *ldADerefOffsetC) String() string { return "LD A, (0xff00 + C)" }
func (l *ldADerefOffsetC) Execute(c *Cpu) uint8 {
	c.ldAOffsetC()
	return 8
}

type ldSPHL struct{}

func (l *ldSPHL) String() string { return "LD SP, HL" }
func (l *ldSPHL) Execute(c *Cpu) uint8 {
	c.ldSPHL()
	return 8
}

type ldADerefIm struct{}

func (l *ldADerefIm) String() string { return "LD A, (u16)" }
func (l *ldADerefIm) Execute(c *Cpu) uint8 {
	c.ldADerefIm()
	return 16
}
