package cpu

type loadByteTarget interface {
	write(byte)
}

type loadByteValue interface {
	value() byte
}

type loadWordTarget interface {
	write(word)
}

type loadWordValue interface {
	value() word
}

func loadWord(from loadWordValue, to loadWordTarget) {
	to.write(from.value())
}

func loadByte(from loadByteValue, to loadByteTarget) {
	to.write(from.value())
}

type doubleRegister struct {
	r [2]*byte
}

func (d *doubleRegister) write(w word) {
	*(d.r[0]) = byte(w >> 8)
	*(d.r[1]) = byte(w & 0xff)
}

type readWord struct {
	w word
}

func (r *readWord) value() word {
	return r.w
}

type readRegister struct {
	r *byte
}

func (r *readRegister) value() byte {
	return *r.r
}

type writeRegister struct {
	r *byte
}

func (w *writeRegister) write(b byte) {
	*w.r = b
}

type readByte struct {
	b byte
}

func (r *readByte) value() byte {
	return r.b
}
