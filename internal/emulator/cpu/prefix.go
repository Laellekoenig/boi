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
