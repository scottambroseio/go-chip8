package display

import "testing"

func TestGetBitsFromByte(t *testing.T) {
	b := byte(0xF0)
	e := [8]byte{1, 1, 1, 1, 0, 0, 0, 0}

	if res := getBitsFromByte(b); res != e {
		t.Errorf("Expected: %v Got: %v", e, res)
	}
}
