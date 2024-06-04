package cpu

type nop struct{}

func (n *nop) String() string { return "NOP" }
func (n *nop) Execute(c *Cpu) uint8 {
	return 4
}

type stop struct{}

func (s *stop) String() string { return "STOP" }
func (s *stop) Execute(c *Cpu) uint8 {
	panic("STOP instruction not implemented")
	// return 4
}

type halt struct{}

func (h *halt) String() string { return "HALT" }
func (h *halt) Execute(c *Cpu) uint8 {
	panic("HALT instruction not implemented")
	//return 4
}

type ei struct{}

func (e *ei) String() string { return "EI" }
func (e *ei) Execute(c *Cpu) uint8 {
	c.ime = true
	return 4
}

type di struct{}

func (d *di) String() string { return "DI" }
func (d *di) Execute(c *Cpu) uint8 {
	c.ime = false
	return 4
}
