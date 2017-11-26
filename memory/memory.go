package memory

// Memory is a 4kb array of bytes representing the memory of the CHIP-8
type Memory [4096]byte

func (m *Memory) WriteBytesAt(b []byte, off int) {
	for i, v := range b {
		m[i+off] = v
	}
}

func (m *Memory) ReadBytesAt(b []byte, off int) {
	for i, v := range m[off : off+len(b)] {
		b[i] = v
	}
}

func (m *Memory) WriteByteAt(b byte, off int) {
	m[off] = b
}

func (m *Memory) ReadByteAt(off int) byte {
	return m[off]
}
