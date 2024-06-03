package cpu

type instruction interface {
	String() string
	Execute(c *Cpu) uint8
}

type notImplemented struct{}

func (n *notImplemented) String() string       { return "Not implemented" }
func (n *notImplemented) Execute(c *Cpu) uint8 { panic("instruction not implemented") }

type nop struct{}

func (n *nop) String() string { return "NOP" }
func (n *nop) Execute(c *Cpu) uint8 {
	return 4
}

type jpIm struct{}

func (j *jpIm) String() string { return "JP u16" }
func (j *jpIm) Execute(c *Cpu) uint8 {
	addr := c.currentWord()
	c.jump(addr, None)
	return 16
}

func InstrucionFromByte(opcode byte) instruction {
	switch opcode {
	case 0x00:
		return &nop{}
	case 0x01:
		return &ldBCIm{}
	case 0x02:
		return &ldDerefBCA{}
	case 0x06:
		return &ldBIm{}
	case 0x08:
		return &ldDerefImSP{}
	case 0x0a:
		return &ldADerefBC{}
	case 0x0e:
		return &ldCIm{}
	case 0x11:
		return &ldDEIm{}
	case 0x12:
		return &ldDerefDEA{}
	case 0x16:
		return &ldDIm{}
	case 0x1a:
		return &ldADerefDE{}
	case 0x1e:
		return &ldEIm{}
	case 0x21:
		return &ldHLIm{}
	case 0x22:
		return &ldDerefIncHLA{}
	case 0x26:
		return &ldHIm{}
	case 0x2a:
		return &ldADerefIncHL{}
	case 0x2e:
		return &ldLIm{}
	case 0x31:
		return &ldSPIm{}
	case 0x32:
		return &ldDerefDecHLA{}
	case 0x36:
		return &ldDerefHLIm{}
	case 0x3a:
		return &ldADerefDecHL{}
	case 0x3e:
		return &ldAIm{}
	case 0x40:
		return &ldBB{}
	case 0x41:
		return &ldBC{}
	case 0x42:
		return &ldBD{}
	case 0x43:
		return &ldBE{}
	case 0x44:
		return &ldBH{}
	case 0x45:
		return &ldBL{}
	case 0x46:
		return &ldBDerefHL{}
	case 0x47:
		return &ldBA{}
	case 0x48:
		return &ldCB{}
	case 0x49:
		return &ldCC{}
	case 0x4a:
		return &ldCD{}
	case 0x4b:
		return &ldCE{}
	case 0x4c:
		return &ldCH{}
	case 0x4d:
		return &ldCL{}
	case 0x4e:
		return &ldCDerefHL{}
	case 0x4f:
		return &ldCA{}
	case 0x50:
		return &ldDB{}
	case 0x51:
		return &ldDC{}
	case 0x52:
		return &ldDD{}
	case 0x53:
		return &ldDE{}
	case 0x54:
		return &ldDH{}
	case 0x55:
		return &ldDL{}
	case 0x56:
		return &ldDDerefHL{}
	case 0x57:
		return &ldDA{}
	case 0x58:
		return &ldEB{}
	case 0x59:
		return &ldEC{}
	case 0x5a:
		return &ldED{}
	case 0x5b:
		return &ldEE{}
	case 0x5c:
		return &ldEH{}
	case 0x5d:
		return &ldEL{}
	case 0x5e:
		return &ldEDerefHL{}
	case 0x5f:
		return &ldEA{}
	case 0x60:
		return &ldHB{}
	case 0x61:
		return &ldHC{}
	case 0x62:
		return &ldHD{}
	case 0x63:
		return &ldHE{}
	case 0x64:
		return &ldHH{}
	case 0x65:
		return &ldHL{}
	case 0x66:
		return &ldHDerefHL{}
	case 0x67:
		return &ldHA{}
	case 0x68:
		return &ldLB{}
	case 0x69:
		return &ldLC{}
	case 0x6a:
		return &ldLD{}
	case 0x6b:
		return &ldLE{}
	case 0x6c:
		return &ldLH{}
	case 0x6d:
		return &ldLL{}
	case 0x6e:
		return &ldLDerefHL{}
	case 0x6f:
		return &ldLA{}
	case 0x70:
		return &ldDerefHLB{}
	case 0x71:
		return &ldDerefHLC{}
	case 0x72:
		return &ldDerefHLD{}
	case 0x73:
		return &ldDerefHLE{}
	case 0x74:
		return &ldDerefHLH{}
	case 0x75:
		return &ldDerefHLL{}
	case 0x77:
		return &ldDerefHLA{}
	case 0x78:
		return &ldAB{}
	case 0x79:
		return &ldAC{}
	case 0x7a:
		return &ldAD{}
	case 0x7b:
		return &ldAE{}
	case 0x7c:
		return &ldAH{}
	case 0x7d:
		return &ldAL{}
	case 0x7e:
		return &ldADerefHL{}
	case 0x7f:
		return &ldAA{}
	case 0xe0:
		return &ldDerefOffsetImA{}
	case 0xe2:
		return &ldDerefOffsetCA{}
	case 0xea:
		return &ldDerefImA{}
	case 0xf0:
		return &ldADerefOffsetIm{}
	case 0xf2:
		return &ldADerefOffsetC{}
	case 0xf9:
		return &ldSPHL{}
	case 0xfa:
		return &ldADerefIm{}
	case 0xc3:
		return &jpIm{}
	}

	return &notImplemented{}
}
