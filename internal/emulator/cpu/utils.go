package cpu

func doesByteAddHalfCarry(a, b byte) bool {
	return (a&0xf)+(b&0xf) > 0xf
}

func doesByteSubHalfCarry(a, b byte) bool {
	return (a & 0xf) < (b & 0xf)
}

func doesWordAddHalfCarry(a, b word) bool {
	return (a&0xfff)+(b&0xfff) > 0xfff
}

func doesWordSubHalfCarry(a, b word) bool {
	return (a & 0xfff) < (b & 0xfff)
}

func doesWordAddCarry(a, b word) bool {
	return uint32(a)+uint32(b) > 0xffff
}

func doesByteAddCarry(a, b byte) bool {
	return uint16(a)+uint16(b) > 0xff
}

func doesByteSubCarry(a, b byte) bool {
	return a < b
}

func doesWordSubCarry(a, b word) bool {
	return a < b
}
