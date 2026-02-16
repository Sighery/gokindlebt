package gokindlebt

/*
#include <kindlebt/kindlebt.h>
#include <kindlebt/kindlebt_log.h>
#include "utils.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type SessionType int

const (
	SessionBle     SessionType = C.ACEBT_SESSION_TYPE_BLE
	SessionDual    SessionType = C.ACEBT_SESSION_TYPE_DUAL_MODE
	SessionClassic SessionType = C.ACEBT_SESSION_TYPE_CLASSIC
	SessionNone    SessionType = C.ACEBT_SESSION_TYPE_NONE
)

type ConnParameter int

const (
	ParameterMax      ConnParameter = C.ACE_BT_BLE_CONN_PARAM_MAX
	ParameterHigh     ConnParameter = C.ACE_BT_BLE_CONN_PARAM_HIGH
	ParameterBalanced ConnParameter = C.ACE_BT_BLE_CONN_PARAM_BALANCED
	ParameterLow      ConnParameter = C.ACE_BT_BLE_CONN_PARAM_LOW
	ParameterUltraLow ConnParameter = C.ACE_BT_BLE_CONN_PARAM_ULTRA_LOW
)

type ConnPriority int

const (
	PriorityLow       ConnPriority = C.ACE_BT_BLE_CONN_PRIO_LOW
	PriorityMedium    ConnPriority = C.ACE_BT_BLE_CONN_PRIO_MEDIUM
	PriorityHigh      ConnPriority = C.ACE_BT_BLE_CONN_PRIO_HIGH
	PriorityDedicated ConnPriority = C.ACE_BT_BLE_CONN_PRIO_DEDICATED
)

type ConnRole int

const (
	RoleClient ConnRole = C.ACEBT_BLE_GATT_CLIENT_ROLE
	RoleSocket ConnRole = C.ACEBT_BLE_SOCKET
)

type RadioState int

const (
	StateDisabled  RadioState = C.ACEBT_STATE_DISABLED
	StateEnabled   RadioState = C.ACEBT_STATE_ENABLED
	StateEnabling  RadioState = C.ACEBT_STATE_ENABLING
	StateDisabling RadioState = C.ACEBT_STATE_DISABLING
)

type Session struct {
	ptr *C.sessionHandle
}

func (session Session) String() string {
	return fmt.Sprintf("%#x", uintptr(unsafe.Pointer(*session.ptr)))
}

type BleConnection struct {
	ptr *C.bleConnHandle
}

func (conn BleConnection) String() string {
	return fmt.Sprintf("%#x", uintptr(unsafe.Pointer(*conn.ptr)))
}

type GattService struct {
	ptr   *C.bleGattsService_t
	noSvc uint
}

func (service GattService) DumpServices() string {
	var size C.size_t = 1024
	var offset C.size_t = 0

	logBuff := (*C.char)(C.malloc(size))
	if logBuff == nil {
		fmt.Printf("Error allocating print buffer for GattService")
		return ""
	}

	svcSlice := unsafe.Slice(service.ptr, service.noSvc)

	for i := C.uint32_t(0); i < C.uint32_t(service.noSvc); i++ {
		fmtStr := C.CString("GATT Database index :%u %p\n")

		logBuff = C.append_to_buffer_wrapper(
			logBuff, &size, &offset, fmtStr, i, unsafe.Pointer(&svcSlice[i]),
		)

		C.free(unsafe.Pointer(fmtStr))

		logBuff = C.utilsDumpServer(
			&svcSlice[i], logBuff, &size, &offset,
		)
	}

	goBytes := C.GoBytes(unsafe.Pointer(logBuff), C.int(offset))
	C.free(unsafe.Pointer(logBuff))
	return string(goBytes)
}

type Notification struct {
	C    <-chan CharacteristicValue
	Stop func() error
}
