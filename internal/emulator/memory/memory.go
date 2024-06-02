package memory

type word = uint16

const (
	romBank0Start = 0x0000
	romBank0End   = 0x3ff
)

type Memory struct {
	data [0xffff + 1]byte
}

func NewMemory(rom []byte) *Memory {
	m := Memory{}

	copy(m.data[:romBank0End+1], rom[:romBank0End+1])

	return &m
}

func (m *Memory) ByteAt(addr word) byte {
	return m.data[addr]
}

func (m *Memory) WordAt(addr word) word {
	panic("implement me")
}
