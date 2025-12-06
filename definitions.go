package gokindlebt

/*
#include <kindlebt/kindlebt.h>
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
	ptr *C.bleGattsService_t
}

type Notification struct {
	C    <-chan CharacteristicValue
	Stop func() error
}
