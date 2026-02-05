package gokindlebt

import (
	"fmt"
	"testing"
)

func TestNewCharacteristicUuidFromString(t *testing.T) {
	testCases := []string{
		"ff120000000000000000000000000000", "FF120000-0000-0000-0000-000000000000",
		"FF120000000000000000000000000000", "ff120000-0000-0000-0000-000000000000",
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			addr, err := NewCharacteristicUuidFromString(tc)
			if err != nil {
				t.Error(err)
			}
			expected := [16]byte{
				0xFF, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00,
			}
			if addr.Bytes != expected {
				t.Errorf("Not expected bytes %v", addr)
			}
		})
	}
}

func TestNewCharacteristicUuidFromStringInvalids(t *testing.T) {
	testCases := []string{
		"", "FF12", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"XXXXXXXXXXXX-XXXX-XXXX-XXXXXXXXXXXX", "XXXXXXXX-XXXX-XXXX-XXXX--XXXXXXXXXXXX",
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc), func(t *testing.T) {
			_, err := NewCharacteristicUuidFromString(tc)
			if err == nil {
				t.Errorf("Value %s got passed as valid", tc)
			}
		})
	}
}

func TestCharacteristicUuidString(t *testing.T) {
	uuid := CharacteristicUuid{Bytes: [16]byte{
		0xFF, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}}
	expected := "ff120000-0000-0000-0000-000000000000"
	if result := uuid.String(); result != expected {
		t.Errorf("Invalid result, got: %s, wanted: %s", result, expected)
	}
}
