package cpu

type Instruction interface {
	String() string
	Timing() uint8
	Execute(c *Cpu)
}

type NotImplemented struct{}

func (n *NotImplemented) String() string { return "Not implemented" }
func (n *NotImplemented) Timing() uint8  { panic("instruction not implemented") }
func (n *NotImplemented) Execute(c *Cpu) { panic("instruction not implemented") }

type NOP struct{}

func (n *NOP) String() string { return "NOP" }
func (n *NOP) Timing() uint8  { return 4 }
func (n *NOP) Execute(c *Cpu) { c.pc += 1 }

func InstrucionFromByte(opcode byte) Instruction {
	switch opcode {
	case 0x00:
		return &NOP{}
	}

	return &NotImplemented{}
}
