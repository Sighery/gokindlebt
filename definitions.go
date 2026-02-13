package gokindlebt

/*
#include <kindlebt/kindlebt.h>
#include "setup.h"
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

func NewSession() (Session, error) {
	ptr := C.newSessionHandle()
	if ptr == nil {
		return Session{}, fmt.Errorf("Failed to allocate session struct")
	}
	return Session{ptr: &ptr}, nil
}

func (session Session) String() string {
	return fmt.Sprintf("%#x", uintptr(unsafe.Pointer(*session.ptr)))
}

func (session Session) Close() {
	// Seems like closeSession will free the memory?
	session.ptr = nil
}

type BleConnection struct {
	ptr *C.bleConnHandle
}

func NewBleConnection() (BleConnection, error) {
	ptr := C.newBleConnHandle()
	if ptr == nil {
		return BleConnection{}, fmt.Errorf("Failed to allocate ble connection struct")
	}
	return BleConnection{ptr: &ptr}, nil
}

func (conn BleConnection) String() string {
	return fmt.Sprintf("%#x", uintptr(unsafe.Pointer(*conn.ptr)))
}

func (conn BleConnection) Close() {
	// Seems like bleDisconnect will free the memory?
	conn.ptr = nil
}

type GattService struct {
	ptr   *C.bleGattsService_t
	noSvc uint
}

func NewGattService() (GattService, error) {
	ptr := C.newBleGattsService()
	if ptr == nil {
		return GattService{}, fmt.Errorf("Failed to allocate gatt db struct")
	}
	return GattService{ptr: ptr, noSvc: 0}, nil
}

func (service GattService) Close() {
	C.free(unsafe.Pointer(service.ptr))
	service.ptr = nil
	service.noSvc = 0
}

type Notification struct {
	C    <-chan CharacteristicValue
	Stop func() error
}
