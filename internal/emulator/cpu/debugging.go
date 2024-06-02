package cpu

import (
	"bytes"
	"fmt"
)

func (c *Cpu) Pc() uint16 {
	return c.pc
}

func (c *Cpu) GetPc() string {
	return fmt.Sprintf("%04x", c.pc)
}

func (c *Cpu) GetSp() string {
	return fmt.Sprintf("%04x", c.sp)
}

func (c *Cpu) GetRegister(t registerType) string {
	var val byte

	switch t {
	case RegA:
		val = c.a
	case RegF:
		val = c.f
	case RegB:
		val = c.b
	case RegC:
		val = c.c
	case RegD:
		val = c.d
	case RegE:
		val = c.e
	case RegH:
		val = c.h
	case RegL:
		val = c.l
	}

	return fmt.Sprintf("%02x", val)
}

func (c *Cpu) GetFlags() string {
	var buf bytes.Buffer

	if c.CheckFlag(FlagZ) {
		buf.WriteString("Z")
	} else {
		buf.WriteString("-")
	}

	if c.CheckFlag(FlagN) {
		buf.WriteString("N")
	} else {
		buf.WriteString("-")
	}

	if c.CheckFlag(FlagH) {
		buf.WriteString("H")
	} else {
		buf.WriteString("-")
	}

	if c.CheckFlag(FlagC) {
		buf.WriteString("C")
	} else {
		buf.WriteString("-")
	}

	return buf.String()
}
