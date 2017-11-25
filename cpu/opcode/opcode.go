package opcode

// Opcode is a 2 byte CHIP-8 opcode
type Opcode uint16

// NewOpcode creates an opcode from 2 consecutive bytes
func NewOpcode(b1, b2 byte) Opcode {
	left := uint16(b1) << 8
	right := uint16(b2)

	return Opcode(left | right)
}

// X returns the second half of the first byte of the opcode
func (o Opcode) X() byte {
	return byte((o & 0x0F00) >> 8)

}

// Y returns the first half of the second byte of the opcode
func (o Opcode) Y() byte {
	return byte((o & 0x00F0) >> 4)
}

// N returns the second half of the second byte of the opcode
func (o Opcode) N() byte {
	return byte(o & 0x000F)
}

// KK returns the second byte of the opcode
func (o Opcode) KK() byte {
	return byte(o & 0x00FF)
}

// NNN returns the result of using the bitwise AND operation on the second half
// of the first byte and the second byte of the opcode
func (o Opcode) NNN() uint16 {
	return uint16(o & 0x0FFF)
}

// LeadByte returns the first half byte of the opcode
func (o Opcode) LeadByte() byte {
	return byte((o & 0xF000) >> 12)
}
