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

type CharacteristicValueKind interface {
	isCharacteristicValueKind()
}

type CharacteristicValueUint8 struct{ V uint8 }

func (CharacteristicValueUint8) isCharacteristicValueKind() {}

type CharacteristicValueUint16 struct{ V uint16 }

func (CharacteristicValueUint16) isCharacteristicValueKind() {}

type CharacteristicValueUint32 struct{ V uint32 }

func (CharacteristicValueUint32) isCharacteristicValueKind() {}

type CharacteristicValueInt8 struct{ V int8 }

func (CharacteristicValueInt8) isCharacteristicValueKind() {}

type CharacteristicValueInt16 struct{ V int16 }

func (CharacteristicValueInt16) isCharacteristicValueKind() {}

type CharacteristicValueInt32 struct{ V int32 }

func (CharacteristicValueInt32) isCharacteristicValueKind() {}

type CharacteristicValueBlob struct{ V []byte }

func (CharacteristicValueBlob) isCharacteristicValueKind() {}

type CharacteristicValue struct {
	Value CharacteristicValueKind
}

func newCharacteristicValueFromC(source C.bleGattCharacteristicsValue_t) CharacteristicValue {
	var value CharacteristicValueKind

	// Anonymous unions seem like a pain to use from Golang
	base := unsafe.Pointer(&source)

	switch source.format {
	case C.BLE_FORMAT_UINT8:
		value = CharacteristicValueUint8{V: *(*uint8)(base)}
	case C.BLE_FORMAT_UINT16:
		value = CharacteristicValueUint16{V: *(*uint16)(base)}
	case C.BLE_FORMAT_UINT32:
		value = CharacteristicValueUint32{V: *(*uint32)(base)}
	case C.BLE_FORMAT_SINT8:
		value = CharacteristicValueInt8{V: *(*int8)(base)}
	case C.BLE_FORMAT_SINT16:
		value = CharacteristicValueInt16{V: *(*int16)(base)}
	case C.BLE_FORMAT_SINT32:
		value = CharacteristicValueInt32{V: *(*int32)(base)}
	case C.BLE_FORMAT_BLOB:
		blob := *(*C.bleGattBlobValue_t)(base)
		ptr := unsafe.Pointer(blob.data)
		if blob.offset > 0 {
			ptr = unsafe.Pointer(uintptr(ptr) + uintptr(blob.offset))
		}
		value = CharacteristicValueBlob{V: C.GoBytes(ptr, C.int(blob.size))}
	}

	return CharacteristicValue{Value: value}
}

type CharacteristicUuid struct {
	Bytes [16]byte
}

func (uuid CharacteristicUuid) String() string {
	return fmt.Sprintf(
		"%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		uuid.Bytes[0], uuid.Bytes[1], uuid.Bytes[2], uuid.Bytes[3], uuid.Bytes[4], uuid.Bytes[5],
		uuid.Bytes[6], uuid.Bytes[7], uuid.Bytes[8], uuid.Bytes[9], uuid.Bytes[10],
		uuid.Bytes[11], uuid.Bytes[12], uuid.Bytes[13], uuid.Bytes[14], uuid.Bytes[15],
	)
}

func NewCharacteristicUuidFromString(uuid string) (CharacteristicUuid, error) {
	if len(uuid) != 32 && len(uuid) != 36 {
		return CharacteristicUuid{}, fmt.Errorf(
			"Use XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX (32 characters) or " +
				"XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX format",
		)
	} else if strings.Contains(uuid, "-") && len(uuid) != 36 {
		return CharacteristicUuid{}, fmt.Errorf("Use XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX format")
	}

	byteAddr, err := hex.DecodeString(strings.ReplaceAll(uuid, "-", ""))
	if err != nil {
		return CharacteristicUuid{}, err
	}

	if byteLen := len(byteAddr); byteLen != 16 {
		return CharacteristicUuid{}, fmt.Errorf("Expected 16 bytes, got %d", byteLen)
	}

	var arr [16]byte
	copy(arr[:], byteAddr)

	return CharacteristicUuid{Bytes: arr}, nil
}

func (uuid *CharacteristicUuid) cUuid() C.uuid_t {
	var cUuid C.uuid_t

	C.memcpy(
		unsafe.Pointer(&cUuid.uu[0]), unsafe.Pointer(&uuid.Bytes[0]),
		C.size_t(len(uuid.Bytes)),
	)
	cUuid._type = C.ACEBT_UUID_TYPE_128

	return cUuid
}
