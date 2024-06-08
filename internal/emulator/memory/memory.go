package memory

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type word = uint16

const (
	romBank0Start = 0x0000
	romBank0End   = 0x3fff
	romBank1Start = 0x4000
	romBank1End   = 0x7fff
)

type Memory struct {
	data      [0xffff + 1]byte
	rom       []byte
	serialOut bytes.Buffer
}

func NewMemory(rom []byte) *Memory {
	m := Memory{
		rom: rom,
	}
	m.mapMemory()

	return &m
}

func (m *Memory) mapMemory() {
	copy(m.data[:romBank0End+1], m.rom[:romBank0End+1])
	copy(m.data[romBank1Start:romBank1End+1], m.rom[romBank0End+1:])

	/*
	  ff00: cf 00 7e ff bd 00 00 f8 ff ff ff ff ff ff ff e1
	  ff10: 80 bf f3 ff bf ff 3f 00 ff bf 7f ff 9f ff bf ff
	  ff20: ff 00 00 bf 77 f3 f1 ff ff ff ff ff ff ff ff ff
	  ff30: bf 08 de 00 bf 11 ff 00 fa 24 ff 10 fd 45 fd 00
	  ff40: 91 81 00 00 90 00 ff fc ff ff 00 00 ff ff ff ff
	  ff50: ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff
	*/

	m.data[0xff00] = 0xcf // P1
	m.data[0xff01] = 0x00 // SB
	m.data[0xff02] = 0x7e // SC
	m.data[0xff03] = 0xff // DIV
	m.data[0xff04] = 0xbd // TIMA
	m.data[0xff05] = 0x00 // TMA
	m.data[0xff06] = 0x00 // TAC
	m.data[0xff07] = 0xf8 // IF
	m.data[0xff08] = 0xff // NR10
	m.data[0xff09] = 0xff // NR11
	m.data[0xff0a] = 0xff // NR12
	m.data[0xff0b] = 0xff // NR13
	m.data[0xff0c] = 0xff // NR14
	m.data[0xff0d] = 0xff // NR21
	m.data[0xff0e] = 0xff // NR22
	m.data[0xff0f] = 0xe1 // NR23

	m.data[0xff10] = 0x80 // NR24
	m.data[0xff11] = 0xbf // NR30
	m.data[0xff12] = 0xf3 // NR31
	m.data[0xff13] = 0xff // NR32
	m.data[0xff14] = 0xbf // NR33
	m.data[0xff15] = 0xff // NR34
	m.data[0xff16] = 0x3f // NR41
	m.data[0xff17] = 0x00 // NR42
	m.data[0xff18] = 0xff // NR43
	m.data[0xff19] = 0xbf // NR44
	m.data[0xff1a] = 0x7f // NR50
	m.data[0xff1b] = 0xff // NR51
	m.data[0xff1c] = 0x9f // NR52
	m.data[0xff1d] = 0xff // LCDC
	m.data[0xff1e] = 0xbf // STAT
	m.data[0xff1f] = 0xff // SCY

	m.data[0xff20] = 0xff // SCX
	m.data[0xff21] = 0x00 // LY
	m.data[0xff22] = 0x00 // LYC
	m.data[0xff23] = 0xbf // DMA
	m.data[0xff24] = 0x77 // BGP
	m.data[0xff25] = 0xf3 // OBP0
	m.data[0xff26] = 0xf1 // OBP1
	m.data[0xff27] = 0xff // WY
	m.data[0xff28] = 0xff // WX
	m.data[0xff29] = 0xff // -
	m.data[0xff2a] = 0xff // -
	m.data[0xff2b] = 0xff // -
	m.data[0xff2c] = 0xff // -
	m.data[0xff2d] = 0xff // -
	m.data[0xff2e] = 0xff // -
	m.data[0xff2f] = 0xff // -

	m.data[0xff30] = 0xbf // -
	m.data[0xff31] = 0x08 // -
	m.data[0xff32] = 0xde // -
	m.data[0xff33] = 0x00 // -
	m.data[0xff34] = 0xbf // -
	m.data[0xff35] = 0x11 // -
	m.data[0xff36] = 0xff // -
	m.data[0xff37] = 0x00 // -
	m.data[0xff38] = 0xfa // -
	m.data[0xff39] = 0x24 // -
	m.data[0xff3a] = 0xff // -
	m.data[0xff3b] = 0x10 // -
	m.data[0xff3c] = 0xfd // -
	m.data[0xff3d] = 0x45 // -
	m.data[0xff3e] = 0xfd // -
	m.data[0xff3f] = 0x00 // -

	m.data[0xff40] = 0x91 // LCDC
	m.data[0xff41] = 0x81 // STAT
	m.data[0xff42] = 0x00 // SCY
	m.data[0xff43] = 0x00 // SCX
	m.data[0xff44] = 0x90 // LY
	m.data[0xff45] = 0x00 // LYC
	m.data[0xff46] = 0xff // DMA
	m.data[0xff47] = 0xfc // BGP
	m.data[0xff48] = 0xff // OBP0
	m.data[0xff49] = 0xff // OBP1
	m.data[0xff4a] = 0x00 // WY
	m.data[0xff4b] = 0x00 // WX
	m.data[0xff4c] = 0xff // -
	m.data[0xff4d] = 0xff // -
	m.data[0xff4e] = 0xff // -
	m.data[0xff50] = 0xff // -
}

func (m *Memory) ByteAt(addr word) byte {
	return m.data[addr]
}

func (m *Memory) WordAt(addr word) word {
	return binary.LittleEndian.Uint16(m.data[addr : addr+2])
}

func (m *Memory) WriteByteAt(b byte, addr word) {
	if addr == 0xff01 {
		m.serialOut.WriteByte(b)
		fmt.Printf("%c", b)
	}

	m.data[addr] = b
}

func (m *Memory) WriteWordAt(w, addr word) {
	binary.LittleEndian.PutUint16(m.data[addr:addr+2], w)
}

func (m *Memory) Reset() {
	m.mapMemory()
}

func (m *Memory) GetSerial() string {
	return m.serialOut.String()
}
