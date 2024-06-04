package cpu

type ei struct{}

func (e *ei) String() string { return "EI" }
func (e *ei) Execute(c *Cpu) uint8 {
	c.ime = true
	return 4
}
