package memory

import (
	"fmt"

	"github.com/Laellekoenig/boi/internal/emulator/cpu"
)

func (m *Memory) GetMemory(loc uint16) []string {
	var lst []string

	var i uint16
	var prev byte
	for i = range 10 {
		b := m.ByteAt(loc + i)
		nb := m.ByteAt(loc + i + 1)

		if prev == 0xcb {
			lst = append(lst, fmt.Sprintf("%04x %02x", loc+i, b))
		} else {
			instruction := cpu.InstrucionFromByte(b, nb)
			lst = append(lst, fmt.Sprintf("%04x %02x %s", loc+i, b, instruction.String()))
		}

		prev = b
	}

	return lst
}
