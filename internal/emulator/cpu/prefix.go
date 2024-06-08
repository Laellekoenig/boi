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
