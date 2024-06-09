package cpu

type rlca struct{}

func (r *rlca) String() string { return "RLCA" }
func (r *rlca) Execute(c *Cpu) uint8 {
	return c.rlca()
}

type rrca struct{}

func (r *rrca) String() string { return "RRCA" }
func (r *rrca) Execute(c *Cpu) uint8 {
	return c.rrca()
}

type rla struct{}

func (r *rla) String() string { return "RLA" }
func (r *rla) Execute(c *Cpu) uint8 {
	return c.rla()
}

type rra struct{}

func (r *rra) String() string { return "RRA" }
func (r *rra) Execute(c *Cpu) uint8 {
	return c.rra()
}

type srlB struct{}

func (s *srlB) String() string { return "SRL B" }
func (s *srlB) Execute(c *Cpu) uint8 {
	return c.srlReg(RegB)
}

type srlC struct{}

func (s *srlC) String() string { return "SRL C" }
func (s *srlC) Execute(c *Cpu) uint8 {
	return c.srlReg(RegC)
}

type srlD struct{}

func (s *srlD) String() string { return "SRL D" }
func (s *srlD) Execute(c *Cpu) uint8 {
	return c.srlReg(RegD)
}

type srlE struct{}

func (s *srlE) String() string { return "SRL E" }
func (s *srlE) Execute(c *Cpu) uint8 {
	return c.srlReg(RegE)
}

type srlH struct{}

func (s *srlH) String() string { return "SRL H" }
func (s *srlH) Execute(c *Cpu) uint8 {
	return c.srlReg(RegH)
}

type srlL struct{}

func (s *srlL) String() string { return "SRL L" }
func (s *srlL) Execute(c *Cpu) uint8 {
	return c.srlReg(RegL)
}

type srlDerefHL struct{}

func (s *srlDerefHL) String() string { return "SRL (HL)" }
func (s *srlDerefHL) Execute(c *Cpu) uint8 {
	return c.srlDerefHL()
}

type srlA struct{}

func (s *srlA) String() string { return "SRL A" }
func (s *srlA) Execute(c *Cpu) uint8 {
	return c.srlReg(RegA)
}

type rrA struct{}

func (r *rrA) String() string { return "RR A" }
func (r *rrA) Execute(c *Cpu) uint8 {
	return c.rrReg(RegA)
}

type rrB struct{}

func (r *rrB) String() string { return "RR B" }
func (r *rrB) Execute(c *Cpu) uint8 {
	return c.rrReg(RegB)
}

type rrC struct{}

func (r *rrC) String() string { return "RR C" }
func (r *rrC) Execute(c *Cpu) uint8 {
	return c.rrReg(RegC)
}

type rrD struct{}

func (r *rrD) String() string { return "RR D" }
func (r *rrD) Execute(c *Cpu) uint8 {
	return c.rrReg(RegD)
}

type rrE struct{}

func (r *rrE) String() string { return "RR E" }
func (r *rrE) Execute(c *Cpu) uint8 {
	return c.rrReg(RegE)
}

type rrH struct{}

func (r *rrH) String() string { return "RR H" }
func (r *rrH) Execute(c *Cpu) uint8 {
	return c.rrReg(RegH)
}

type rrL struct{}

func (r *rrL) String() string { return "RR L" }
func (r *rrL) Execute(c *Cpu) uint8 {
	return c.rrReg(RegL)
}

type rrDerefHL struct{}

func (r *rrDerefHL) String() string { return "RR (HL)" }
func (r *rrDerefHL) Execute(c *Cpu) uint8 {
	return c.rrDerefHL()
}

type rlcB struct{}

func (r *rlcB) String() string { return "RLC B" }
func (r *rlcB) Execute(c *Cpu) uint8 {
	return c.rlcReg(RegB)
}

type rlcC struct{}

func (r *rlcC) String() string { return "RLC C" }
func (r *rlcC) Execute(c *Cpu) uint8 {
	return c.rlcReg(RegC)
}

type rlcD struct{}

func (r *rlcD) String() string { return "RLC D" }
func (r *rlcD) Execute(c *Cpu) uint8 {
	return c.rlcReg(RegD)
}

type rlcE struct{}

func (r *rlcE) String() string { return "RLC E" }
func (r *rlcE) Execute(c *Cpu) uint8 {
	return c.rlcReg(RegE)
}

type rlcH struct{}

func (r *rlcH) String() string { return "RLC H" }
func (r *rlcH) Execute(c *Cpu) uint8 {
	return c.rlcReg(RegH)
}

type rlcL struct{}

func (r *rlcL) String() string { return "RLC L" }
func (r *rlcL) Execute(c *Cpu) uint8 {
	return c.rlcReg(RegL)
}

type rlcDerefHL struct{}

func (r *rlcDerefHL) String() string { return "RLC (HL)" }
func (r *rlcDerefHL) Execute(c *Cpu) uint8 {
	return c.rlcDerefHL()
}

type rlcA struct{}

func (r *rlcA) String() string { return "RLC A" }
func (r *rlcA) Execute(c *Cpu) uint8 {
	return c.rlcReg(RegA)
}

type rrcB struct{}

func (r *rrcB) String() string { return "RRC B" }
func (r *rrcB) Execute(c *Cpu) uint8 {
	return c.rrcReg(RegB)
}

type rrcC struct{}

func (r *rrcC) String() string { return "RRC C" }
func (r *rrcC) Execute(c *Cpu) uint8 {
	return c.rrcReg(RegC)
}

type rrcD struct{}

func (r *rrcD) String() string { return "RRC D" }
func (r *rrcD) Execute(c *Cpu) uint8 {
	return c.rrcReg(RegD)
}

type rrcE struct{}

func (r *rrcE) String() string { return "RRC E" }
func (r *rrcE) Execute(c *Cpu) uint8 {
	return c.rrcReg(RegE)
}

type rrcH struct{}

func (r *rrcH) String() string { return "RRC H" }
func (r *rrcH) Execute(c *Cpu) uint8 {
	return c.rrcReg(RegH)
}

type rrcL struct{}

func (r *rrcL) String() string { return "RRC L" }
func (r *rrcL) Execute(c *Cpu) uint8 {
	return c.rrcReg(RegL)
}

type rrcDerefHL struct{}

func (r *rrcDerefHL) String() string { return "RRC (HL)" }
func (r *rrcDerefHL) Execute(c *Cpu) uint8 {
	return c.rrcDerefHL()
}

type rrcA struct{}

func (r *rrcA) String() string { return "RRC A" }
func (r *rrcA) Execute(c *Cpu) uint8 {
	return c.rrcReg(RegA)
}

type rlB struct{}

func (r *rlB) String() string { return "RL B" }
func (r *rlB) Execute(c *Cpu) uint8 {
	return c.rlReg(RegB)
}

type rlC struct{}

func (r *rlC) String() string { return "RL C" }
func (r *rlC) Execute(c *Cpu) uint8 {
	return c.rlReg(RegC)
}

type rlD struct{}

func (r *rlD) String() string { return "RL D" }
func (r *rlD) Execute(c *Cpu) uint8 {
	return c.rlReg(RegD)
}

type rlE struct{}

func (r *rlE) String() string { return "RL E" }
func (r *rlE) Execute(c *Cpu) uint8 {
	return c.rlReg(RegE)
}

type rlH struct{}

func (r *rlH) String() string { return "RL H" }
func (r *rlH) Execute(c *Cpu) uint8 {
	return c.rlReg(RegH)
}

type rlL struct{}

func (r *rlL) String() string { return "RL L" }
func (r *rlL) Execute(c *Cpu) uint8 {
	return c.rlReg(RegL)
}

type rlDerefHL struct{}

func (r *rlDerefHL) String() string { return "RL (HL)" }
func (r *rlDerefHL) Execute(c *Cpu) uint8 {
	return c.rlDerefHL()
}

type rlA struct{}

func (r *rlA) String() string { return "RL A" }
func (r *rlA) Execute(c *Cpu) uint8 {
	return c.rlReg(RegA)
}

type slaB struct{}

func (s *slaB) String() string { return "SLA B" }
func (s *slaB) Execute(c *Cpu) uint8 {
	return c.slaReg(RegB)
}

type slaC struct{}

func (s *slaC) String() string { return "SLA C" }
func (s *slaC) Execute(c *Cpu) uint8 {
	return c.slaReg(RegC)
}

type slaD struct{}

func (s *slaD) String() string { return "SLA D" }
func (s *slaD) Execute(c *Cpu) uint8 {
	return c.slaReg(RegD)
}

type slaE struct{}

func (s *slaE) String() string { return "SLA E" }
func (s *slaE) Execute(c *Cpu) uint8 {
	return c.slaReg(RegE)
}

type slaH struct{}

func (s *slaH) String() string { return "SLA H" }
func (s *slaH) Execute(c *Cpu) uint8 {
	return c.slaReg(RegH)
}

type slaL struct{}

func (s *slaL) String() string { return "SLA L" }
func (s *slaL) Execute(c *Cpu) uint8 {
	return c.slaReg(RegL)
}

type slaDerefHL struct{}

func (s *slaDerefHL) String() string { return "SLA (HL)" }
func (s *slaDerefHL) Execute(c *Cpu) uint8 {
	return c.slaDerefHL()
}

type slaA struct{}

func (s *slaA) String() string { return "SLA A" }
func (s *slaA) Execute(c *Cpu) uint8 {
	return c.slaReg(RegA)
}

type sraB struct{}

func (s *sraB) String() string { return "SRA B" }
func (s *sraB) Execute(c *Cpu) uint8 {
	return c.sraReg(RegB)
}

type sraC struct{}

func (s *sraC) String() string { return "SRA C" }
func (s *sraC) Execute(c *Cpu) uint8 {
	return c.sraReg(RegC)
}

type sraD struct{}

func (s *sraD) String() string { return "SRA D" }
func (s *sraD) Execute(c *Cpu) uint8 {
	return c.sraReg(RegD)
}

type sraE struct{}

func (s *sraE) String() string { return "SRA E" }
func (s *sraE) Execute(c *Cpu) uint8 {
	return c.sraReg(RegE)
}

type sraH struct{}

func (s *sraH) String() string { return "SRA H" }
func (s *sraH) Execute(c *Cpu) uint8 {
	return c.sraReg(RegH)
}

type sraL struct{}

func (s *sraL) String() string { return "SRA L" }
func (s *sraL) Execute(c *Cpu) uint8 {
	return c.sraReg(RegL)
}

type sraDerefHL struct{}

func (s *sraDerefHL) String() string { return "SRA (HL)" }
func (s *sraDerefHL) Execute(c *Cpu) uint8 {
	return c.sraDerefHL()
}

type sraA struct{}

func (s *sraA) String() string { return "SRA A" }
func (s *sraA) Execute(c *Cpu) uint8 {
	return c.sraReg(RegA)
}

type swapB struct{}

func (s *swapB) String() string { return "SWAP B" }
func (s *swapB) Execute(c *Cpu) uint8 {
	return c.swapReg(RegB)
}

type swapC struct{}

func (s *swapC) String() string { return "SWAP C" }
func (s *swapC) Execute(c *Cpu) uint8 {
	return c.swapReg(RegC)
}

type swapD struct{}

func (s *swapD) String() string { return "SWAP D" }
func (s *swapD) Execute(c *Cpu) uint8 {
	return c.swapReg(RegD)
}

type swapE struct{}

func (s *swapE) String() string { return "SWAP E" }
func (s *swapE) Execute(c *Cpu) uint8 {
	return c.swapReg(RegE)
}

type swapH struct{}

func (s *swapH) String() string { return "SWAP H" }
func (s *swapH) Execute(c *Cpu) uint8 {
	return c.swapReg(RegH)
}

type swapL struct{}

func (s *swapL) String() string { return "SWAP L" }
func (s *swapL) Execute(c *Cpu) uint8 {
	return c.swapReg(RegL)
}

type swapDerefHL struct{}

func (s *swapDerefHL) String() string { return "SWAP (HL)" }
func (s *swapDerefHL) Execute(c *Cpu) uint8 {
	return c.swapDerefHL()
}

type swapA struct{}

func (s *swapA) String() string { return "SWAP A" }
func (s *swapA) Execute(c *Cpu) uint8 {
	return c.swapReg(RegA)
}

type bit0B struct{}

func (b *bit0B) String() string { return "BIT 0, B" }
func (b *bit0B) Execute(c *Cpu) uint8 {
	return c.bitReg(0, RegB)
}

type bit0C struct{}

func (b *bit0C) String() string { return "BIT 0, C" }
func (b *bit0C) Execute(c *Cpu) uint8 {
	return c.bitReg(0, RegC)
}

type bit0D struct{}

func (b *bit0D) String() string { return "BIT 0, D" }
func (b *bit0D) Execute(c *Cpu) uint8 {
	return c.bitReg(0, RegD)
}

type bit0E struct{}

func (b *bit0E) String() string { return "BIT 0, E" }
func (b *bit0E) Execute(c *Cpu) uint8 {
	return c.bitReg(0, RegE)
}

type bit0H struct{}

func (b *bit0H) String() string { return "BIT 0, H" }
func (b *bit0H) Execute(c *Cpu) uint8 {
	return c.bitReg(0, RegH)
}

type bit0L struct{}

func (b *bit0L) String() string { return "BIT 0, L" }
func (b *bit0L) Execute(c *Cpu) uint8 {
	return c.bitReg(0, RegL)
}

type bit0DerefHL struct{}

func (b *bit0DerefHL) String() string { return "BIT 0, (HL)" }
func (b *bit0DerefHL) Execute(c *Cpu) uint8 {
	return c.bitDerefHL(0)
}

type bit0A struct{}

func (b *bit0A) String() string { return "BIT 0, A" }
func (b *bit0A) Execute(c *Cpu) uint8 {
	return c.bitReg(0, RegA)
}

type bit1B struct{}

func (b *bit1B) String() string { return "BIT 1, B" }
func (b *bit1B) Execute(c *Cpu) uint8 {
	return c.bitReg(1, RegB)
}

type bit1C struct{}

func (b *bit1C) String() string { return "BIT 1, C" }
func (b *bit1C) Execute(c *Cpu) uint8 {
	return c.bitReg(1, RegC)
}

type bit1D struct{}

func (b *bit1D) String() string { return "BIT 1, D" }
func (b *bit1D) Execute(c *Cpu) uint8 {
	return c.bitReg(1, RegD)
}

type bit1E struct{}

func (b *bit1E) String() string { return "BIT 1, E" }
func (b *bit1E) Execute(c *Cpu) uint8 {
	return c.bitReg(1, RegE)
}

type bit1H struct{}

func (b *bit1H) String() string { return "BIT 1, H" }
func (b *bit1H) Execute(c *Cpu) uint8 {
	return c.bitReg(1, RegH)
}

type bit1L struct{}

func (b *bit1L) String() string { return "BIT 1, L" }
func (b *bit1L) Execute(c *Cpu) uint8 {
	return c.bitReg(1, RegL)
}

type bit1DerefHL struct{}

func (b *bit1DerefHL) String() string { return "BIT 1, (HL)" }
func (b *bit1DerefHL) Execute(c *Cpu) uint8 {
	return c.bitDerefHL(1)
}

type bit1A struct{}

func (b *bit1A) String() string { return "BIT 1, A" }
func (b *bit1A) Execute(c *Cpu) uint8 {
	return c.bitReg(1, RegA)
}

type bit2B struct{}

func (b *bit2B) String() string { return "BIT 2, B" }
func (b *bit2B) Execute(c *Cpu) uint8 {
	return c.bitReg(2, RegB)
}

type bit2C struct{}

func (b *bit2C) String() string { return "BIT 2, C" }
func (b *bit2C) Execute(c *Cpu) uint8 {
	return c.bitReg(2, RegC)
}

type bit2D struct{}

func (b *bit2D) String() string { return "BIT 2, D" }
func (b *bit2D) Execute(c *Cpu) uint8 {
	return c.bitReg(2, RegD)
}

type bit2E struct{}

func (b *bit2E) String() string { return "BIT 2, E" }
func (b *bit2E) Execute(c *Cpu) uint8 {
	return c.bitReg(2, RegE)
}

type bit2H struct{}

func (b *bit2H) String() string { return "BIT 2, H" }
func (b *bit2H) Execute(c *Cpu) uint8 {
	return c.bitReg(2, RegH)
}

type bit2L struct{}

func (b *bit2L) String() string { return "BIT 2, L" }
func (b *bit2L) Execute(c *Cpu) uint8 {
	return c.bitReg(2, RegL)
}

type bit2DerefHL struct{}

func (b *bit2DerefHL) String() string { return "BIT 2, (HL)" }
func (b *bit2DerefHL) Execute(c *Cpu) uint8 {
	return c.bitDerefHL(2)
}

type bit2A struct{}

func (b *bit2A) String() string { return "BIT 2, A" }
func (b *bit2A) Execute(c *Cpu) uint8 {
	return c.bitReg(2, RegA)
}

type bit3B struct{}

func (b *bit3B) String() string { return "BIT 3, B" }
func (b *bit3B) Execute(c *Cpu) uint8 {
	return c.bitReg(3, RegB)
}

type bit3C struct{}

func (b *bit3C) String() string { return "BIT 3, C" }
func (b *bit3C) Execute(c *Cpu) uint8 {
	return c.bitReg(3, RegC)
}

type bit3D struct{}

func (b *bit3D) String() string { return "BIT 3, D" }
func (b *bit3D) Execute(c *Cpu) uint8 {
	return c.bitReg(3, RegD)
}

type bit3E struct{}

func (b *bit3E) String() string { return "BIT 3, E" }
func (b *bit3E) Execute(c *Cpu) uint8 {
	return c.bitReg(3, RegE)
}

type bit3H struct{}

func (b *bit3H) String() string { return "BIT 3, H" }
func (b *bit3H) Execute(c *Cpu) uint8 {
	return c.bitReg(3, RegH)
}

type bit3L struct{}

func (b *bit3L) String() string { return "BIT 3, L" }
func (b *bit3L) Execute(c *Cpu) uint8 {
	return c.bitReg(3, RegL)
}

type bit3DerefHL struct{}

func (b *bit3DerefHL) String() string { return "BIT 3, (HL)" }
func (b *bit3DerefHL) Execute(c *Cpu) uint8 {
	return c.bitDerefHL(3)
}

type bit3A struct{}

func (b *bit3A) String() string { return "BIT 3, A" }
func (b *bit3A) Execute(c *Cpu) uint8 {
	return c.bitReg(3, RegA)
}

type bit4B struct{}

func (b *bit4B) String() string { return "BIT 4, B" }
func (b *bit4B) Execute(c *Cpu) uint8 {
	return c.bitReg(4, RegB)
}

type bit4C struct{}

func (b *bit4C) String() string { return "BIT 4, C" }
func (b *bit4C) Execute(c *Cpu) uint8 {
	return c.bitReg(4, RegC)
}

type bit4D struct{}

func (b *bit4D) String() string { return "BIT 4, D" }
func (b *bit4D) Execute(c *Cpu) uint8 {
	return c.bitReg(4, RegD)
}

type bit4E struct{}

func (b *bit4E) String() string { return "BIT 4, E" }
func (b *bit4E) Execute(c *Cpu) uint8 {
	return c.bitReg(4, RegE)
}

type bit4H struct{}

func (b *bit4H) String() string { return "BIT 4, H" }
func (b *bit4H) Execute(c *Cpu) uint8 {
	return c.bitReg(4, RegH)
}

type bit4L struct{}

func (b *bit4L) String() string { return "BIT 4, L" }
func (b *bit4L) Execute(c *Cpu) uint8 {
	return c.bitReg(4, RegL)
}

type bit4DerefHL struct{}

func (b *bit4DerefHL) String() string { return "BIT 4, (HL)" }
func (b *bit4DerefHL) Execute(c *Cpu) uint8 {
	return c.bitDerefHL(4)
}

type bit4A struct{}

func (b *bit4A) String() string { return "BIT 4, A" }
func (b *bit4A) Execute(c *Cpu) uint8 {
	return c.bitReg(4, RegA)
}

type bit5B struct{}

func (b *bit5B) String() string { return "BIT 5, B" }
func (b *bit5B) Execute(c *Cpu) uint8 {
	return c.bitReg(5, RegB)
}

type bit5C struct{}

func (b *bit5C) String() string { return "BIT 5, C" }
func (b *bit5C) Execute(c *Cpu) uint8 {
	return c.bitReg(5, RegC)
}

type bit5D struct{}

func (b *bit5D) String() string { return "BIT 5, D" }
func (b *bit5D) Execute(c *Cpu) uint8 {
	return c.bitReg(5, RegD)
}

type bit5E struct{}

func (b *bit5E) String() string { return "BIT 5, E" }
func (b *bit5E) Execute(c *Cpu) uint8 {
	return c.bitReg(5, RegE)
}

type bit5H struct{}

func (b *bit5H) String() string { return "BIT 5, H" }
func (b *bit5H) Execute(c *Cpu) uint8 {
	return c.bitReg(5, RegH)
}

type bit5L struct{}

func (b *bit5L) String() string { return "BIT 5, L" }
func (b *bit5L) Execute(c *Cpu) uint8 {
	return c.bitReg(5, RegL)
}

type bit5DerefHL struct{}

func (b *bit5DerefHL) String() string { return "BIT 5, (HL)" }
func (b *bit5DerefHL) Execute(c *Cpu) uint8 {
	return c.bitDerefHL(5)
}

type bit5A struct{}

func (b *bit5A) String() string { return "BIT 5, A" }
func (b *bit5A) Execute(c *Cpu) uint8 {
	return c.bitReg(5, RegA)
}

type bit6B struct{}

func (b *bit6B) String() string { return "BIT 6, B" }
func (b *bit6B) Execute(c *Cpu) uint8 {
	return c.bitReg(6, RegB)
}

type bit6C struct{}

func (b *bit6C) String() string { return "BIT 6, C" }
func (b *bit6C) Execute(c *Cpu) uint8 {
	return c.bitReg(6, RegC)
}

type bit6D struct{}

func (b *bit6D) String() string { return "BIT 6, D" }
func (b *bit6D) Execute(c *Cpu) uint8 {
	return c.bitReg(6, RegD)
}

type bit6E struct{}

func (b *bit6E) String() string { return "BIT 6, E" }
func (b *bit6E) Execute(c *Cpu) uint8 {
	return c.bitReg(6, RegE)
}

type bit6H struct{}

func (b *bit6H) String() string { return "BIT 6, H" }
func (b *bit6H) Execute(c *Cpu) uint8 {
	return c.bitReg(6, RegH)
}

type bit6L struct{}

func (b *bit6L) String() string { return "BIT 6, L" }
func (b *bit6L) Execute(c *Cpu) uint8 {
	return c.bitReg(6, RegL)
}

type bit6DerefHL struct{}

func (b *bit6DerefHL) String() string { return "BIT 6, (HL)" }
func (b *bit6DerefHL) Execute(c *Cpu) uint8 {
	return c.bitDerefHL(6)
}

type bit6A struct{}

func (b *bit6A) String() string { return "BIT 6, A" }
func (b *bit6A) Execute(c *Cpu) uint8 {
	return c.bitReg(6, RegA)
}

type bit7B struct{}

func (b *bit7B) String() string { return "BIT 7, B" }
func (b *bit7B) Execute(c *Cpu) uint8 {
	return c.bitReg(7, RegB)
}

type bit7C struct{}

func (b *bit7C) String() string { return "BIT 7, C" }
func (b *bit7C) Execute(c *Cpu) uint8 {
	return c.bitReg(7, RegC)
}

type bit7D struct{}

func (b *bit7D) String() string { return "BIT 7, D" }
func (b *bit7D) Execute(c *Cpu) uint8 {
	return c.bitReg(7, RegD)
}

type bit7E struct{}

func (b *bit7E) String() string { return "BIT 7, E" }
func (b *bit7E) Execute(c *Cpu) uint8 {
	return c.bitReg(7, RegE)
}

type bit7H struct{}

func (b *bit7H) String() string { return "BIT 7, H" }
func (b *bit7H) Execute(c *Cpu) uint8 {
	return c.bitReg(7, RegH)
}

type bit7L struct{}

func (b *bit7L) String() string { return "BIT 7, L" }
func (b *bit7L) Execute(c *Cpu) uint8 {
	return c.bitReg(7, RegL)
}

type bit7DerefHL struct{}

func (b *bit7DerefHL) String() string { return "BIT 7, (HL)" }
func (b *bit7DerefHL) Execute(c *Cpu) uint8 {
	return c.bitDerefHL(7)
}

type bit7A struct{}

func (b *bit7A) String() string { return "BIT 7, A" }
func (b *bit7A) Execute(c *Cpu) uint8 {
	return c.bitReg(7, RegA)
}

type res0B struct{}

func (r *res0B) String() string { return "RES 0, B" }
func (r *res0B) Execute(c *Cpu) uint8 {
	return c.resReg(0, RegB)
}

type res0C struct{}

func (r *res0C) String() string { return "RES 0, C" }
func (r *res0C) Execute(c *Cpu) uint8 {
	return c.resReg(0, RegC)
}

type res0D struct{}

func (r *res0D) String() string { return "RES 0, D" }
func (r *res0D) Execute(c *Cpu) uint8 {
	return c.resReg(0, RegD)
}

type res0E struct{}

func (r *res0E) String() string { return "RES 0, E" }
func (r *res0E) Execute(c *Cpu) uint8 {
	return c.resReg(0, RegE)
}

type res0H struct{}

func (r *res0H) String() string { return "RES 0, H" }
func (r *res0H) Execute(c *Cpu) uint8 {
	return c.resReg(0, RegH)
}

type res0L struct{}

func (r *res0L) String() string { return "RES 0, L" }
func (r *res0L) Execute(c *Cpu) uint8 {
	return c.resReg(0, RegL)
}

type res0DerefHL struct{}

func (r *res0DerefHL) String() string { return "RES 0, (HL)" }
func (r *res0DerefHL) Execute(c *Cpu) uint8 {
	return c.resDerefHL(0)
}

type res0A struct{}

func (r *res0A) String() string { return "RES 0, A" }
func (r *res0A) Execute(c *Cpu) uint8 {
	return c.resReg(0, RegA)
}

type res1B struct{}

func (r *res1B) String() string { return "RES 1, B" }
func (r *res1B) Execute(c *Cpu) uint8 {
	return c.resReg(1, RegB)
}

type res1C struct{}

func (r *res1C) String() string { return "RES 1, C" }
func (r *res1C) Execute(c *Cpu) uint8 {
	return c.resReg(1, RegC)
}

type res1D struct{}

func (r *res1D) String() string { return "RES 1, D" }
func (r *res1D) Execute(c *Cpu) uint8 {
	return c.resReg(1, RegD)
}

type res1E struct{}

func (r *res1E) String() string { return "RES 1, E" }
func (r *res1E) Execute(c *Cpu) uint8 {
	return c.resReg(1, RegE)
}

type res1H struct{}

func (r *res1H) String() string { return "RES 1, H" }
func (r *res1H) Execute(c *Cpu) uint8 {
	return c.resReg(1, RegH)
}

type res1L struct{}

func (r *res1L) String() string { return "RES 1, L" }
func (r *res1L) Execute(c *Cpu) uint8 {
	return c.resReg(1, RegL)
}

type res1DerefHL struct{}

func (r *res1DerefHL) String() string { return "RES 1, (HL)" }
func (r *res1DerefHL) Execute(c *Cpu) uint8 {
	return c.resDerefHL(1)
}

type res1A struct{}

func (r *res1A) String() string { return "RES 1, A" }
func (r *res1A) Execute(c *Cpu) uint8 {
	return c.resReg(1, RegA)
}

type res2B struct{}

func (r *res2B) String() string { return "RES 2, B" }
func (r *res2B) Execute(c *Cpu) uint8 {
	return c.resReg(2, RegB)
}

type res2C struct{}

func (r *res2C) String() string { return "RES 2, C" }
func (r *res2C) Execute(c *Cpu) uint8 {
	return c.resReg(2, RegC)
}

type res2D struct{}

func (r *res2D) String() string { return "RES 2, D" }
func (r *res2D) Execute(c *Cpu) uint8 {
	return c.resReg(2, RegD)
}

type res2E struct{}

func (r *res2E) String() string { return "RES 2, E" }
func (r *res2E) Execute(c *Cpu) uint8 {
	return c.resReg(2, RegE)
}

type res2H struct{}

func (r *res2H) String() string { return "RES 2, H" }
func (r *res2H) Execute(c *Cpu) uint8 {
	return c.resReg(2, RegH)
}

type res2L struct{}

func (r *res2L) String() string { return "RES 2, L" }
func (r *res2L) Execute(c *Cpu) uint8 {
	return c.resReg(2, RegL)
}

type res2DerefHL struct{}

func (r *res2DerefHL) String() string { return "RES 2, (HL)" }
func (r *res2DerefHL) Execute(c *Cpu) uint8 {
	return c.resDerefHL(2)
}

type res2A struct{}

func (r *res2A) String() string { return "RES 2, A" }
func (r *res2A) Execute(c *Cpu) uint8 {
	return c.resReg(2, RegA)
}

type res3B struct{}

func (r *res3B) String() string { return "RES 3, B" }
func (r *res3B) Execute(c *Cpu) uint8 {
	return c.resReg(3, RegB)
}

type res3C struct{}

func (r *res3C) String() string { return "RES 3, C" }
func (r *res3C) Execute(c *Cpu) uint8 {
	return c.resReg(3, RegC)
}

type res3D struct{}

func (r *res3D) String() string { return "RES 3, D" }
func (r *res3D) Execute(c *Cpu) uint8 {
	return c.resReg(3, RegD)
}

type res3E struct{}

func (r *res3E) String() string { return "RES 3, E" }
func (r *res3E) Execute(c *Cpu) uint8 {
	return c.resReg(3, RegE)
}

type res3H struct{}

func (r *res3H) String() string { return "RES 3, H" }
func (r *res3H) Execute(c *Cpu) uint8 {
	return c.resReg(3, RegH)
}

type res3L struct{}

func (r *res3L) String() string { return "RES 3, L" }
func (r *res3L) Execute(c *Cpu) uint8 {
	return c.resReg(3, RegL)
}

type res3DerefHL struct{}

func (r *res3DerefHL) String() string { return "RES 3, (HL)" }
func (r *res3DerefHL) Execute(c *Cpu) uint8 {
	return c.resDerefHL(3)
}

type res3A struct{}

func (r *res3A) String() string { return "RES 3, A" }
func (r *res3A) Execute(c *Cpu) uint8 {
	return c.resReg(3, RegA)
}

type res4B struct{}

func (r *res4B) String() string { return "RES 4, B" }
func (r *res4B) Execute(c *Cpu) uint8 {
	return c.resReg(4, RegB)
}

type res4C struct{}

func (r *res4C) String() string { return "RES 4, C" }
func (r *res4C) Execute(c *Cpu) uint8 {
	return c.resReg(4, RegC)
}

type res4D struct{}

func (r *res4D) String() string { return "RES 4, D" }
func (r *res4D) Execute(c *Cpu) uint8 {
	return c.resReg(4, RegD)
}

type res4E struct{}

func (r *res4E) String() string { return "RES 4, E" }
func (r *res4E) Execute(c *Cpu) uint8 {
	return c.resReg(4, RegE)
}

type res4H struct{}

func (r *res4H) String() string { return "RES 4, H" }
func (r *res4H) Execute(c *Cpu) uint8 {
	return c.resReg(4, RegH)
}

type res4L struct{}

func (r *res4L) String() string { return "RES 4, L" }
func (r *res4L) Execute(c *Cpu) uint8 {
	return c.resReg(4, RegL)
}

type res4DerefHL struct{}

func (r *res4DerefHL) String() string { return "RES 4, (HL)" }
func (r *res4DerefHL) Execute(c *Cpu) uint8 {
	return c.resDerefHL(4)
}

type res4A struct{}

func (r *res4A) String() string { return "RES 4, A" }
func (r *res4A) Execute(c *Cpu) uint8 {
	return c.resReg(4, RegA)
}

type res5B struct{}

func (r *res5B) String() string { return "RES 5, B" }
func (r *res5B) Execute(c *Cpu) uint8 {
	return c.resReg(5, RegB)
}

type res5C struct{}

func (r *res5C) String() string { return "RES 5, C" }
func (r *res5C) Execute(c *Cpu) uint8 {
	return c.resReg(5, RegC)
}

type res5D struct{}

func (r *res5D) String() string { return "RES 5, D" }
func (r *res5D) Execute(c *Cpu) uint8 {
	return c.resReg(5, RegD)
}

type res5E struct{}

func (r *res5E) String() string { return "RES 5, E" }
func (r *res5E) Execute(c *Cpu) uint8 {
	return c.resReg(5, RegE)
}

type res5H struct{}

func (r *res5H) String() string { return "RES 5, H" }
func (r *res5H) Execute(c *Cpu) uint8 {
	return c.resReg(5, RegH)
}

type res5L struct{}

func (r *res5L) String() string { return "RES 5, L" }
func (r *res5L) Execute(c *Cpu) uint8 {
	return c.resReg(5, RegL)
}

type res5DerefHL struct{}

func (r *res5DerefHL) String() string { return "RES 5, (HL)" }
func (r *res5DerefHL) Execute(c *Cpu) uint8 {
	return c.resDerefHL(5)
}

type res5A struct{}

func (r *res5A) String() string { return "RES 5, A" }
func (r *res5A) Execute(c *Cpu) uint8 {
	return c.resReg(5, RegA)
}

type res6B struct{}

func (r *res6B) String() string { return "RES 6, B" }
func (r *res6B) Execute(c *Cpu) uint8 {
	return c.resReg(6, RegB)
}

type res6C struct{}

func (r *res6C) String() string { return "RES 6, C" }
func (r *res6C) Execute(c *Cpu) uint8 {
	return c.resReg(6, RegC)
}

type res6D struct{}

func (r *res6D) String() string { return "RES 6, D" }
func (r *res6D) Execute(c *Cpu) uint8 {
	return c.resReg(6, RegD)
}

type res6E struct{}

func (r *res6E) String() string { return "RES 6, E" }
func (r *res6E) Execute(c *Cpu) uint8 {
	return c.resReg(6, RegE)
}

type res6H struct{}

func (r *res6H) String() string { return "RES 6, H" }
func (r *res6H) Execute(c *Cpu) uint8 {
	return c.resReg(6, RegH)
}

type res6L struct{}

func (r *res6L) String() string { return "RES 6, L" }
func (r *res6L) Execute(c *Cpu) uint8 {
	return c.resReg(6, RegL)
}

type res6DerefHL struct{}

func (r *res6DerefHL) String() string { return "RES 6, (HL)" }
func (r *res6DerefHL) Execute(c *Cpu) uint8 {
	return c.resDerefHL(6)
}

type res6A struct{}

func (r *res6A) String() string { return "RES 6, A" }
func (r *res6A) Execute(c *Cpu) uint8 {
	return c.resReg(6, RegA)
}

type res7B struct{}

func (r *res7B) String() string { return "RES 7, B" }
func (r *res7B) Execute(c *Cpu) uint8 {
	return c.resReg(7, RegB)
}

type res7C struct{}

func (r *res7C) String() string { return "RES 7, C" }
func (r *res7C) Execute(c *Cpu) uint8 {
	return c.resReg(7, RegC)
}

type res7D struct{}

func (r *res7D) String() string { return "RES 7, D" }
func (r *res7D) Execute(c *Cpu) uint8 {
	return c.resReg(7, RegD)
}

type res7E struct{}

func (r *res7E) String() string { return "RES 7, E" }
func (r *res7E) Execute(c *Cpu) uint8 {
	return c.resReg(7, RegE)
}

type res7H struct{}

func (r *res7H) String() string { return "RES 7, H" }
func (r *res7H) Execute(c *Cpu) uint8 {
	return c.resReg(7, RegH)
}

type res7L struct{}

func (r *res7L) String() string { return "RES 7, L" }
func (r *res7L) Execute(c *Cpu) uint8 {
	return c.resReg(7, RegL)
}

type res7DerefHL struct{}

func (r *res7DerefHL) String() string { return "RES 7, (HL)" }
func (r *res7DerefHL) Execute(c *Cpu) uint8 {
	return c.resDerefHL(7)
}

type res7A struct{}

func (r *res7A) String() string { return "RES 7, A" }
func (r *res7A) Execute(c *Cpu) uint8 {
	return c.resReg(7, RegA)
}

type set0B struct{}

func (s *set0B) String() string { return "SET 0, B" }
func (s *set0B) Execute(c *Cpu) uint8 {
	return c.setReg(0, RegB)
}

type set0C struct{}

func (s *set0C) String() string { return "SET 0, C" }
func (s *set0C) Execute(c *Cpu) uint8 {
	return c.setReg(0, RegC)
}

type set0D struct{}

func (s *set0D) String() string { return "SET 0, D" }
func (s *set0D) Execute(c *Cpu) uint8 {
	return c.setReg(0, RegD)
}

type set0E struct{}

func (s *set0E) String() string { return "SET 0, E" }
func (s *set0E) Execute(c *Cpu) uint8 {
	return c.setReg(0, RegE)
}

type set0H struct{}

func (s *set0H) String() string { return "SET 0, H" }
func (s *set0H) Execute(c *Cpu) uint8 {
	return c.setReg(0, RegH)
}

type set0L struct{}

func (s *set0L) String() string { return "SET 0, L" }
func (s *set0L) Execute(c *Cpu) uint8 {
	return c.setReg(0, RegL)
}

type set0DerefHL struct{}

func (s *set0DerefHL) String() string { return "SET 0, (HL)" }
func (s *set0DerefHL) Execute(c *Cpu) uint8 {
	return c.setDerefHL(0)
}

type set0A struct{}

func (s *set0A) String() string { return "SET 0, A" }
func (s *set0A) Execute(c *Cpu) uint8 {
	return c.setReg(0, RegA)
}

type set1B struct{}

func (s *set1B) String() string { return "SET 1, B" }
func (s *set1B) Execute(c *Cpu) uint8 {
	return c.setReg(1, RegB)
}

type set1C struct{}

func (s *set1C) String() string { return "SET 1, C" }
func (s *set1C) Execute(c *Cpu) uint8 {
	return c.setReg(1, RegC)
}

type set1D struct{}

func (s *set1D) String() string { return "SET 1, D" }
func (s *set1D) Execute(c *Cpu) uint8 {
	return c.setReg(1, RegD)
}

type set1E struct{}

func (s *set1E) String() string { return "SET 1, E" }
func (s *set1E) Execute(c *Cpu) uint8 {
	return c.setReg(1, RegE)
}

type set1H struct{}

func (s *set1H) String() string { return "SET 1, H" }
func (s *set1H) Execute(c *Cpu) uint8 {
	return c.setReg(1, RegH)
}

type set1L struct{}

func (s *set1L) String() string { return "SET 1, L" }
func (s *set1L) Execute(c *Cpu) uint8 {
	return c.setReg(1, RegL)
}

type set1DerefHL struct{}

func (s *set1DerefHL) String() string { return "SET 1, (HL)" }
func (s *set1DerefHL) Execute(c *Cpu) uint8 {
	return c.setDerefHL(1)
}

type set1A struct{}

func (s *set1A) String() string { return "SET 1, A" }
func (s *set1A) Execute(c *Cpu) uint8 {
	return c.setReg(1, RegA)
}

type set2B struct{}

func (s *set2B) String() string { return "SET 2, B" }
func (s *set2B) Execute(c *Cpu) uint8 {
	return c.setReg(2, RegB)
}

type set2C struct{}

func (s *set2C) String() string { return "SET 2, C" }
func (s *set2C) Execute(c *Cpu) uint8 {
	return c.setReg(2, RegC)
}

type set2D struct{}

func (s *set2D) String() string { return "SET 2, D" }
func (s *set2D) Execute(c *Cpu) uint8 {
	return c.setReg(2, RegD)
}

type set2E struct{}

func (s *set2E) String() string { return "SET 2, E" }
func (s *set2E) Execute(c *Cpu) uint8 {
	return c.setReg(2, RegE)
}

type set2H struct{}

func (s *set2H) String() string { return "SET 2, H" }
func (s *set2H) Execute(c *Cpu) uint8 {
	return c.setReg(2, RegH)
}

type set2L struct{}

func (s *set2L) String() string { return "SET 2, L" }
func (s *set2L) Execute(c *Cpu) uint8 {
	return c.setReg(2, RegL)
}

type set2DerefHL struct{}

func (s *set2DerefHL) String() string { return "SET 2, (HL)" }
func (s *set2DerefHL) Execute(c *Cpu) uint8 {
	return c.setDerefHL(2)
}

type set2A struct{}

func (s *set2A) String() string { return "SET 2, A" }
func (s *set2A) Execute(c *Cpu) uint8 {
	return c.setReg(2, RegA)
}

type set3B struct{}

func (s *set3B) String() string { return "SET 3, B" }
func (s *set3B) Execute(c *Cpu) uint8 {
	return c.setReg(3, RegB)
}

type set3C struct{}

func (s *set3C) String() string { return "SET 3, C" }
func (s *set3C) Execute(c *Cpu) uint8 {
	return c.setReg(3, RegC)
}

type set3D struct{}

func (s *set3D) String() string { return "SET 3, D" }
func (s *set3D) Execute(c *Cpu) uint8 {
	return c.setReg(3, RegD)
}

type set3E struct{}

func (s *set3E) String() string { return "SET 3, E" }
func (s *set3E) Execute(c *Cpu) uint8 {
	return c.setReg(3, RegE)
}

type set3H struct{}

func (s *set3H) String() string { return "SET 3, H" }
func (s *set3H) Execute(c *Cpu) uint8 {
	return c.setReg(3, RegH)
}

type set3L struct{}

func (s *set3L) String() string { return "SET 3, L" }
func (s *set3L) Execute(c *Cpu) uint8 {
	return c.setReg(3, RegL)
}

type set3DerefHL struct{}

func (s *set3DerefHL) String() string { return "SET 3, (HL)" }
func (s *set3DerefHL) Execute(c *Cpu) uint8 {
	return c.setDerefHL(3)
}

type set3A struct{}

func (s *set3A) String() string { return "SET 3, A" }
func (s *set3A) Execute(c *Cpu) uint8 {
	return c.setReg(3, RegA)
}

type set4B struct{}

func (s *set4B) String() string { return "SET 4, B" }
func (s *set4B) Execute(c *Cpu) uint8 {
	return c.setReg(4, RegB)
}

type set4C struct{}

func (s *set4C) String() string { return "SET 4, C" }
func (s *set4C) Execute(c *Cpu) uint8 {
	return c.setReg(4, RegC)
}

type set4D struct{}

func (s *set4D) String() string { return "SET 4, D" }
func (s *set4D) Execute(c *Cpu) uint8 {
	return c.setReg(4, RegD)
}

type set4E struct{}

func (s *set4E) String() string { return "SET 4, E" }
func (s *set4E) Execute(c *Cpu) uint8 {
	return c.setReg(4, RegE)
}

type set4H struct{}

func (s *set4H) String() string { return "SET 4, H" }
func (s *set4H) Execute(c *Cpu) uint8 {
	return c.setReg(4, RegH)
}

type set4L struct{}

func (s *set4L) String() string { return "SET 4, L" }
func (s *set4L) Execute(c *Cpu) uint8 {
	return c.setReg(4, RegL)
}

type set4DerefHL struct{}

func (s *set4DerefHL) String() string { return "SET 4, (HL)" }
func (s *set4DerefHL) Execute(c *Cpu) uint8 {
	return c.setDerefHL(4)
}

type set4A struct{}

func (s *set4A) String() string { return "SET 4, A" }
func (s *set4A) Execute(c *Cpu) uint8 {
	return c.setReg(4, RegA)
}

type set5B struct{}

func (s *set5B) String() string { return "SET 5, B" }
func (s *set5B) Execute(c *Cpu) uint8 {
	return c.setReg(5, RegB)
}

type set5C struct{}

func (s *set5C) String() string { return "SET 5, C" }
func (s *set5C) Execute(c *Cpu) uint8 {
	return c.setReg(5, RegC)
}

type set5D struct{}

func (s *set5D) String() string { return "SET 5, D" }
func (s *set5D) Execute(c *Cpu) uint8 {
	return c.setReg(5, RegD)
}

type set5E struct{}

func (s *set5E) String() string { return "SET 5, E" }
func (s *set5E) Execute(c *Cpu) uint8 {
	return c.setReg(5, RegE)
}

type set5H struct{}

func (s *set5H) String() string { return "SET 5, H" }
func (s *set5H) Execute(c *Cpu) uint8 {
	return c.setReg(5, RegH)
}

type set5L struct{}

func (s *set5L) String() string { return "SET 5, L" }
func (s *set5L) Execute(c *Cpu) uint8 {
	return c.setReg(5, RegL)
}

type set5DerefHL struct{}

func (s *set5DerefHL) String() string { return "SET 5, (HL)" }
func (s *set5DerefHL) Execute(c *Cpu) uint8 {
	return c.setDerefHL(5)
}

type set5A struct{}

func (s *set5A) String() string { return "SET 5, A" }
func (s *set5A) Execute(c *Cpu) uint8 {
	return c.setReg(5, RegA)
}

type set6B struct{}

func (s *set6B) String() string { return "SET 6, B" }
func (s *set6B) Execute(c *Cpu) uint8 {
	return c.setReg(6, RegB)
}

type set6C struct{}

func (s *set6C) String() string { return "SET 6, C" }
func (s *set6C) Execute(c *Cpu) uint8 {
	return c.setReg(6, RegC)
}

type set6D struct{}

func (s *set6D) String() string { return "SET 6, D" }
func (s *set6D) Execute(c *Cpu) uint8 {
	return c.setReg(6, RegD)
}

type set6E struct{}

func (s *set6E) String() string { return "SET 6, E" }
func (s *set6E) Execute(c *Cpu) uint8 {
	return c.setReg(6, RegE)
}

type set6H struct{}

func (s *set6H) String() string { return "SET 6, H" }
func (s *set6H) Execute(c *Cpu) uint8 {
	return c.setReg(6, RegH)
}

type set6L struct{}

func (s *set6L) String() string { return "SET 6, L" }
func (s *set6L) Execute(c *Cpu) uint8 {
	return c.setReg(6, RegL)
}

type set6DerefHL struct{}

func (s *set6DerefHL) String() string { return "SET 6, (HL)" }
func (s *set6DerefHL) Execute(c *Cpu) uint8 {
	return c.setDerefHL(6)
}

type set6A struct{}

func (s *set6A) String() string { return "SET 6, A" }
func (s *set6A) Execute(c *Cpu) uint8 {
	return c.setReg(6, RegA)
}

type set7B struct{}

func (s *set7B) String() string { return "SET 7, B" }
func (s *set7B) Execute(c *Cpu) uint8 {
	return c.setReg(7, RegB)
}

type set7C struct{}

func (s *set7C) String() string { return "SET 7, C" }
func (s *set7C) Execute(c *Cpu) uint8 {
	return c.setReg(7, RegC)
}

type set7D struct{}

func (s *set7D) String() string { return "SET 7, D" }
func (s *set7D) Execute(c *Cpu) uint8 {
	return c.setReg(7, RegD)
}

type set7E struct{}

func (s *set7E) String() string { return "SET 7, E" }
func (s *set7E) Execute(c *Cpu) uint8 {
	return c.setReg(7, RegE)
}

type set7H struct{}

func (s *set7H) String() string { return "SET 7, H" }
func (s *set7H) Execute(c *Cpu) uint8 {
	return c.setReg(7, RegH)
}

type set7L struct{}

func (s *set7L) String() string { return "SET 7, L" }
func (s *set7L) Execute(c *Cpu) uint8 {
	return c.setReg(7, RegL)
}

type set7DerefHL struct{}

func (s *set7DerefHL) String() string { return "SET 7, (HL)" }
func (s *set7DerefHL) Execute(c *Cpu) uint8 {
	return c.setDerefHL(7)
}

type set7A struct{}

func (s *set7A) String() string { return "SET 7, A" }
func (s *set7A) Execute(c *Cpu) uint8 {
	return c.setReg(7, RegA)
}
