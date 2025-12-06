package gokindlebt

import (
	"testing"
)

func TestIsASCIIPrintableValid(t *testing.T) {
	source := []byte{0x79, 0x70, 0x70}
	if result := IsASCIIPrintable(source); result != true {
		t.Errorf("Invalid result, got: %t, wanted: %t", result, true)
	}
}

func TestIsASCIIPrintableInvalid(t *testing.T) {
	source := []byte{0x80, 0x02, 0x70}
	if result := IsASCIIPrintable(source); result != false {
		t.Errorf("Invalid result, got: %t, wanted: %t", result, false)
	}
}
