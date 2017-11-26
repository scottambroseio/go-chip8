package opcode

import "testing"

// add more test data
// rather just just using the same opcode every time
func TestNewOpcode(t *testing.T) {
	e := Opcode(8287)
	if o := NewOpcode(0x20, 0x5F); o != e {
		t.Fatalf("Expected: 0x%04X Got: 0x%04X", e, o)
	}
}
func TestX(t *testing.T) {
	e := byte(0x0)
	if o := Opcode(8287); o.X() != e {
		t.Fatalf("Expected: 0x%04X Got: 0x%04X", e, o)
	}
}

func TestY(t *testing.T) {
	e := byte(0x5)
	if o := Opcode(8287); o.Y() != e {
		t.Fatalf("Expected: 0x%04X Got: 0x%04X", e, o)
	}
}

func TestN(t *testing.T) {
	e := byte(0xF)
	if o := Opcode(8287); o.N() != e {
		t.Fatalf("Expected: 0x%04X Got: 0x%04X", e, o)
	}
}

func TestKK(t *testing.T) {
	e := byte(0x5f)
	if o := Opcode(8287); o.KK() != e {
		t.Fatalf("Expected: 0x%04X Got: 0x%04X", e, o)
	}
}

func TestNNN(t *testing.T) {
	e := uint16(0x5f)
	if o := Opcode(8287); o.NNN() != e {
		t.Fatalf("Expected: 0x%04X Got: 0x%04X", e, o)
	}
}

func TestLeadByte(t *testing.T) {
	e := byte(0x2)
	if o := Opcode(8287); o.LeadByte() != e {
		t.Fatalf("Expected: 0x%04X Got: 0x%04X", e, o)
	}
}
