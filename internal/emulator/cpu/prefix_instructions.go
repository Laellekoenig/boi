package cpu

func PrefixInstrucionFromByte(opcode byte) instruction {
	switch opcode {
	case 0x18:
		return &rrB{}
	case 0x19:
		return &rrC{}
	case 0x1a:
		return &rrD{}
	case 0x1b:
		return &rrE{}
	case 0x1c:
		return &rrH{}
	case 0x1d:
		return &rrL{}
	case 0x1e:
		return &rrDerefHL{}
	case 0x1f:
		return &rrA{}

	case 0x38:
		return &srlB{}
	case 0x39:
		return &srlC{}
	case 0x3a:
		return &srlD{}
	case 0x3b:
		return &srlE{}
	case 0x3c:
		return &srlH{}
	case 0x3d:
		return &srlL{}
	case 0x3e:
		return &srlDerefHL{}
	case 0x3f:
		return &srlA{}

	default:
		return &notImplemented{opcode}
	}
}
