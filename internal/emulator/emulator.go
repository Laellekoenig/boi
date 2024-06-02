package emulator

import (
	"fmt"
	"os"

	"github.com/Laellekoenig/boi/internal/emulator/cpu"
	"github.com/Laellekoenig/boi/internal/emulator/memory"
)

type Emulator struct {
	romFileName string
	Cpu         *cpu.Cpu
	Memory      *memory.Memory
}

func NewEmulator(filename string) *Emulator {
	rom, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read %s: %s\n", filename, err)
		os.Exit(1)
	}

	m := memory.NewMemory(rom)
	e := Emulator{
		romFileName: filename,
		Cpu:         cpu.NewCpu(m),
		Memory:      m,
	}

	return &e
}

func (e *Emulator) RomName() string {
	return e.romFileName
}
