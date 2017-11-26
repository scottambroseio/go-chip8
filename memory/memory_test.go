package memory

import "testing"
import "bytes"

func TestWriteBytesAt_NoOffset(t *testing.T) {
	d := []byte{0x01, 0xC5, 0xE2, 0xF3}
	m := new(Memory)

	m.WriteBytesAt(d, 0)

	if r := m[:len(d)]; !bytes.Equal(r, d) {
		t.Fatalf("Expected to write %v, wrote %v", r, d)
	}
}

func TestWriteBytesAt_WithOffset(t *testing.T) {
	d := []byte{0x01, 0xC5, 0xE2, 0xF3}
	m := new(Memory)
	o := 12
	m.WriteBytesAt(d, o)

	if r := m[o : o+len(d)]; !bytes.Equal(r, d) {
		t.Fatalf("Expected to write %v, wrote %v", r, d)
	}
}

func TestReadBytesAt_NoOffset(t *testing.T) {
	e := []byte{0x01, 0xC5, 0xE2, 0xF3}
	r := make([]byte, 4, 4)
	m := new(Memory)
	m[0] = 0x01
	m[1] = 0xC5
	m[2] = 0xE2
	m[3] = 0xF3

	m.ReadBytesAt(r, 0)

	if r := m[:4]; !bytes.Equal(e, r) {
		t.Fatalf("Expected to read %v, got %v", e, r)
	}
}

func TestReadBytesAt_WithOffset(t *testing.T) {
	e := []byte{0x01, 0xC5, 0xE2, 0xF3}
	r := make([]byte, 4, 4)
	m := new(Memory)
	m[12] = 0x01
	m[13] = 0xC5
	m[14] = 0xE2
	m[15] = 0xF3
	o := 12

	m.ReadBytesAt(r, o)

	if r := m[o : o+4]; !bytes.Equal(e, r) {
		t.Fatalf("Expected to read %v, got %v", e, r)
	}
}

func TestWriteByteAt_NoOffset(t *testing.T) {
	b := byte(0x2F)
	m := new(Memory)
	m.WriteByteAt(b, 0)

	if m[0] != b {
		t.Fatalf("Expected to write %v, wrote %v", b, m[0])
	}
}

func TestWriteByteAt_WithOffset(t *testing.T) {
	o := 12
	b := byte(0x2F)
	m := new(Memory)
	m.WriteByteAt(b, o)

	if m[o] != b {
		t.Fatalf("Expected to write %v, wrote %v", b, m[o])
	}
}

func TestReadByteAt_NoOffset(t *testing.T) {
	b := byte(0x2F)
	m := new(Memory)
	m[0] = b

	if r := m.ReadByteAt(0); r != b {
		t.Fatalf("Expected to read %v, read %v", b, r)
	}
}

func TestReadByteAt_WithOffset(t *testing.T) {
	o := 12
	b := byte(0x2F)
	m := new(Memory)
	m[o] = b

	if r := m.ReadByteAt(o); r != b {
		t.Fatalf("Expected to read %v, read %v", b, r)
	}
}
