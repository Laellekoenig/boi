package memory

import "encoding/binary"

type word = uint16

const (
	romBank0Start = 0x0000
	romBank0End   = 0x3fff
	romBank1Start = 0x4000
	romBank1End   = 0x7fff
)

type Memory struct {
	data [0xffff + 1]byte
}

func NewMemory(rom []byte) *Memory {
	m := Memory{}

	copy(m.data[:romBank0End+1], rom[:romBank0End+1])
	copy(m.data[romBank1Start:romBank1End+1], rom[romBank0End+1:])

	return &m
}

func (m *Memory) ByteAt(addr word) byte {
	return m.data[addr]
}

func (m *Memory) WordAt(addr word) word {
	return binary.LittleEndian.Uint16(m.data[addr : addr+2])
}

func (m *Memory) WriteByteAt(b byte, addr word) {
	m.data[addr] = b
}
