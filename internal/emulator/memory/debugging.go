package memory

import (
	"fmt"

	"github.com/Laellekoenig/boi/internal/emulator/cpu"
)

func (m *Memory) GetMemory(loc uint16) []string {
	var lst []string

	var i uint16
	for i = range 10 {
		b := m.ByteAt(loc + i)
		instruction := cpu.InstrucionFromByte(b)
		lst = append(lst, fmt.Sprintf("%04x %02x %s", loc+i, b, instruction.String()))
	}

	return lst
}
