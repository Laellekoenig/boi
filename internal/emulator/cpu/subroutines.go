package cpu

type jrIm struct{}

func (j *jrIm) String() string { return "JR i8" }
func (j *jrIm) Execute(c *Cpu) uint8 {
	return c.jumpRelativeIm(None)
}

type jrNZIm struct{}

func (j *jrNZIm) String() string { return "JR NZ, i8" }
func (j *jrNZIm) Execute(c *Cpu) uint8 {
	return c.jumpRelativeIm(NZ)
}

type jrZIm struct{}

func (j *jrZIm) String() string { return "JR Z, i8" }
func (j *jrZIm) Execute(c *Cpu) uint8 {
	return c.jumpRelativeIm(Z)
}

type jrNCIm struct{}

func (j *jrNCIm) String() string { return "JR NC, i8" }
func (j *jrNCIm) Execute(c *Cpu) uint8 {
	return c.jumpRelativeIm(NC)
}

type jrCIm struct{}

func (j *jrCIm) String() string { return "JR C, i8" }
func (j *jrCIm) Execute(c *Cpu) uint8 {
	return c.jumpRelativeIm(C)
}

type retNZ struct{}

func (r *retNZ) String() string { return "RET NZ" }
func (r *retNZ) Execute(c *Cpu) uint8 {
	return c.ret(NZ)
}

type jpNZIm struct{}

func (j *jpNZIm) String() string { return "JP NZ, u16" }
func (j *jpNZIm) Execute(c *Cpu) uint8 {
	return c.jumpIm(NZ)
}

type jpIm struct{}

func (j *jpIm) String() string { return "JP u16" }
func (j *jpIm) Execute(c *Cpu) uint8 {
	return c.jumpIm(None)
}

type callNZIm struct{}

func (c *callNZIm) String() string { return "CALL NZ, u16" }
func (c *callNZIm) Execute(cpu *Cpu) uint8 {
	return cpu.callIm(NZ)
}

type rst00h struct{}

func (r *rst00h) String() string { return "RST 00h" }
func (r *rst00h) Execute(cpu *Cpu) uint8 {
	return cpu.rst(0x0000)
}

type retZ struct{}

func (r *retZ) String() string { return "RET Z" }
func (r *retZ) Execute(c *Cpu) uint8 {
	return c.ret(Z)
}

type ret struct{}

func (r *ret) String() string { return "RET" }
func (r *ret) Execute(c *Cpu) uint8 {
	return c.ret(None)
}

type jpZIm struct{}

func (j *jpZIm) String() string { return "JP Z, u16" }
func (j *jpZIm) Execute(c *Cpu) uint8 {
	return c.jumpIm(Z)
}

type callZIm struct{}

func (c *callZIm) String() string { return "CALL Z, u16" }
func (c *callZIm) Execute(cpu *Cpu) uint8 {
	return cpu.callIm(Z)
}

type callIm struct{}

func (c *callIm) String() string { return "CALL u16" }
func (c *callIm) Execute(cpu *Cpu) uint8 {
	return cpu.callIm(None)
}

type rst08h struct{}

func (r *rst08h) String() string { return "RST 08h" }
func (r *rst08h) Execute(cpu *Cpu) uint8 {
	return cpu.rst(0x0008)
}

type retNC struct{}

func (r *retNC) String() string { return "RET NC" }
func (r *retNC) Execute(c *Cpu) uint8 {
	return c.ret(NC)
}

type jpNCIm struct{}

func (j *jpNCIm) String() string { return "JP NC, u16" }
func (j *jpNCIm) Execute(c *Cpu) uint8 {
	return c.jumpIm(NC)
}

type callNCIm struct{}

func (c *callNCIm) String() string { return "CALL NC, u16" }
func (c *callNCIm) Execute(cpu *Cpu) uint8 {
	return cpu.callIm(NC)
}

type rst10h struct{}

func (r *rst10h) String() string { return "RST 10h" }
func (r *rst10h) Execute(cpu *Cpu) uint8 {
	return cpu.rst(0x0010)
}

type retC struct{}

func (r *retC) String() string { return "RET C" }
func (r *retC) Execute(c *Cpu) uint8 {
	return c.ret(C)
}

type reti struct{}

func (r *reti) String() string { return "RETI" }
func (r *reti) Execute(c *Cpu) uint8 {
	return c.reti()
}

type jpCIm struct{}

func (j *jpCIm) String() string { return "JP C, u16" }
func (j *jpCIm) Execute(c *Cpu) uint8 {
	return c.jumpIm(C)
}

type callCIm struct{}

func (c *callCIm) String() string { return "CALL C, u16" }
func (c *callCIm) Execute(cpu *Cpu) uint8 {
	return cpu.callIm(C)
}

type rst18h struct{}

func (r *rst18h) String() string { return "RST 18h" }
func (r *rst18h) Execute(cpu *Cpu) uint8 {
	return cpu.rst(0x0018)
}

type rst20h struct{}

func (r *rst20h) String() string { return "RST 20h" }
func (r *rst20h) Execute(cpu *Cpu) uint8 {
	return cpu.rst(0x0020)
}

type jpHL struct{}

func (j *jpHL) String() string { return "JP HL" }
func (j *jpHL) Execute(c *Cpu) uint8 {
	return c.jpHL()
}

type rst28h struct{}

func (r *rst28h) String() string { return "RST 28h" }
func (r *rst28h) Execute(cpu *Cpu) uint8 {
	return cpu.rst(0x0028)
}

type rst30h struct{}

func (r *rst30h) String() string { return "RST 30h" }
func (r *rst30h) Execute(cpu *Cpu) uint8 {
	return cpu.rst(0x0030)
}

type rst38h struct{}

func (r *rst38h) String() string { return "RST 38h" }
func (r *rst38h) Execute(cpu *Cpu) uint8 {
	return cpu.rst(0x0038)
}
