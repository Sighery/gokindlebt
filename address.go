package gokindlebt

/*
#include <kindlebt/kindlebt.h>
*/
import "C"

import (
	"encoding/hex"
	"fmt"
	"strings"
	"unsafe"
)

type Address struct {
	Bytes [6]byte
}

func NewAddressFromString(addr string) (Address, error) {
	if len(addr) != 17 && len(addr) != 12 {
		return Address{}, fmt.Errorf("Use AA:BB:CC:DD:EE:FF or AABBCCDDEEFF format")
	} else if strings.Contains(addr, ":") && len(addr) != 17 {
		return Address{}, fmt.Errorf("Use AA:BB:CC:DD:EE:FF format")
	}

	byteAddr, err := hex.DecodeString(strings.ReplaceAll(addr, ":", ""))
	if err != nil {
		return Address{}, err
	}

	if byteLen := len(byteAddr); byteLen != 6 {
		return Address{}, fmt.Errorf("Expected 6 bytes, got %d", byteLen)
	}

	var arr [6]byte
	copy(arr[:], byteAddr)

	return Address{Bytes: arr}, nil
}

func (addr *Address) cAddr() *C.bdAddr_t {
	return (*C.bdAddr_t)(unsafe.Pointer(addr))
}

func (addr Address) String() string {
	return fmt.Sprintf(
		"%02X:%02X:%02X:%02X:%02X:%02X",
		addr.Bytes[0], addr.Bytes[1], addr.Bytes[2], addr.Bytes[3],
		addr.Bytes[4], addr.Bytes[5],
	)
}
