package gokindlebt

import (
	"fmt"
	"testing"
)

func TestNewAddressFromStringWithColon(t *testing.T) {
	addr, err := NewAddressFromString("AA:BB:CC:DD:EE:FF")

	if err != nil {
		t.Error(err)
	}

	if addr.Bytes != [6]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF} {
		t.Errorf("Not expected bytes %v", addr)
	}
}

func TestNewAddressFromStringWithoutColon(t *testing.T) {
	addr, err := NewAddressFromString("AABBCCDDEEFF")

	if err != nil {
		t.Error(err)
	}

	if addr.Bytes != [6]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF} {
		t.Errorf("Not expected bytes %v", addr)
	}
}

func TestNewAddressFromStringInvalids(t *testing.T) {
	testCases := []string{
		"", "AABBCCDDEEF", "AABBCCDDEEFFG", "AA:BB:CC:DD:EE:F", "AA:BB:CC:DD:EE:FF:GG",
		"PP:PP:PP:PP:PP:PP", "PPPPPPPPPPPP",
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc), func(t *testing.T) {
			_, err := NewAddressFromString(tc)
			if err == nil {
				t.Errorf("Value %s got passed as valid", tc)
			}
		})
	}
}

func TestAddressString(t *testing.T) {
	addr := Address{Bytes: [6]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}}
	expected := "AA:BB:CC:DD:EE:FF"
	if result := addr.String(); result != expected {
		t.Errorf("Invalid result, got: %s, wanted: %s", result, expected)
	}
}
