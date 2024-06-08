package cpu

func PrefixInstrucionFromByte(opcode byte) instruction {
	switch opcode {
	default:
		return &notImplemented{opcode}
	}
}
