package cpu

const (
	vBlankInterruptV word = 0x0040
	lcdcInterruptV   word = 0x0048
	timerInterruptV  word = 0x0050
	serialInterruptV word = 0x0058
	joypadInterruptV word = 0x0060
)

const (
	interruptEnable word = 0xffff
	interruptFlag   word = 0xff0f
)

func (c *Cpu) handleInterrupts() uint8 {
	if !c.ime {
		return 0
	}

	ifVal := c.bus.ByteAt(interruptFlag)
	ieVal := c.bus.ByteAt(interruptEnable)
	interrupt := ifVal & ieVal

	for i := 0; i <= 4; i++ {
		if interrupt&(1<<i) == 0 {
			continue
		}

		c.ime = false
		updateIf := ifVal & (^(1 << i))
		c.bus.WriteByteAt(updateIf, interruptFlag)

		var target word
		switch i {
		case 0:
			target = vBlankInterruptV
		case 1:
			target = lcdcInterruptV
		case 2:
			target = timerInterruptV
		case 3:
			target = serialInterruptV
		case 4:
			target = joypadInterruptV
		default:
			panic("Invalid interrupt")
		}

		c.pushWord(c.pc)
		c.pc = target

		return 8
	}

	return 0
}
