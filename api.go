package gokindlebt

/*
#include <stdbool.h>

#include <kindlebt/kindlebt.h>
#include <kindlebt/kindlebt_application.h>

#include "callbacks.h"
#include "utils.h"
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

type api interface {
	IsBleSupported() bool
	GetSupportedSession() int
	GetRadioState() (RadioState, error)
	EnableRadio(session Session) error
	DisableRadio(session Session) error
	OpenSession(sessionType int) (Session, error)
	CloseSession(session Session) (Session, error)
	RegisterBle(session Session) error
	DeregisterBle(session Session) error
	RegisterGattClient(session Session) error
	DeregisterGattClient(session Session) error
	ConnectBleSimple(session Session, addr Address) (BleConnection, error)
	ConnectBle(
		session Session, addr Address, param ConnParameter, role ConnRole, priority ConnPriority,
	) (BleConnection, error)
	DisconnectBle(conn BleConnection) error
	DiscoverServices(session Session, conn BleConnection) error
	RetrieveGattDatabase(conn BleConnection) (GattService, error)
	CleanupGattDatabase(service GattService) error
	ReadCharacteristic(
		session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
	) error
	writeCharacteristic(
		session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
		value CharacteristicValueKind, responseRequired C.responseType_t,
	) (CharacteristicValue, error)
	WriteCharacteristicWithoutResponse(
		session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
		value CharacteristicValueKind,
	) error
	WriteCharacteristicWithResponse(
		session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
		value CharacteristicValueKind,
	) (CharacteristicValue, error)
	NotifyCharacteristic(
		session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
	) (*Notification, error)
}

func (a *Adapter) IsBleSupported() bool {
	return bool(C.isBLESupported())
}

func (a *Adapter) GetSupportedSession() int {
	return int(C.getSupportedSession())
}

func (a *Adapter) GetRadioState() (RadioState, error) {
	var state C.state_t
	status := C.getRadioState(&state)
	if status != C.ACE_STATUS_OK {
		return StateEnabled, fmt.Errorf("Couldn't fetch radio state, error: %d", status)
	}
	return RadioState(state), nil
}

func (a *Adapter) EnableRadio(session Session) error {
	status := C.enableRadio(*session.ptr)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Couldn't enable radio, error: %d", status)
	}
	return nil
}

func (a *Adapter) DisableRadio(session Session) error {
	status := C.disableRadio(*session.ptr)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Couldn't disable radio, error: %d", status)
	}
	return nil
}

func (a *Adapter) OpenSession(sessionType SessionType) (Session, error) {
	var session C.sessionHandle

	status := C.openSession((C.sessionType_t)(sessionType), &session)
	if status != C.ACE_STATUS_OK {
		return Session{}, fmt.Errorf("Couldn't open a session, error: %d", status)
	}

	return Session{ptr: &session}, nil
}

func (a *Adapter) CloseSession(session Session) error {
	status := C.closeSession(*session.ptr)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Couldn't close session, error: %d", status)
	}

	return nil
}

func (a *Adapter) RegisterBle(session Session) error {
	status := C.bleRegister(*session.ptr)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Couldn't register for BLE use, error: %d", status)
	}

	return nil
}

func (a *Adapter) DeregisterBle(session Session) error {
	status := C.bleDeregister(*session.ptr)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Couldn't deregister from BLE use, error %d", status)
	}

	return nil
}

func (a *Adapter) RegisterGattClient(session Session) error {
	C.setGcCallbacks(&C.application_gatt_client_callbacks)

	status := C.bleRegisterGattClient(*session.ptr, &C.application_gatt_client_callbacks)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Couldn't register for GATT Client use, error %d", status)
	}

	return nil
}

func (a *Adapter) DeregisterGattClient(session Session) error {
	status := C.bleDeregisterGattClient(*session.ptr)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Couldn't deregister from GATT Client use, error %d", status)
	}

	return nil
}

func (a *Adapter) ConnectBle(
	session Session, addr Address, param ConnParameter, role ConnRole, priority ConnPriority,
) (BleConnection, error) {
	var connPtr C.bleConnHandle

	cAddr := addr.cAddr()

	status := C.bleConnect(
		*session.ptr, &connPtr, cAddr,
		(C.bleConnParam_t)(param),
		(C.bleConnRole_t)(role),
		(C.bleConnPriority_t)(priority),
	)
	if status != C.ACE_STATUS_OK {
		return BleConnection{}, fmt.Errorf(
			"Failed to connect to BLE device %s, error %d", addr, status,
		)
	}

	conn := BleConnection{ptr: &connPtr}
	connRegistry[*conn.ptr] = &conn
	return conn, nil
}

func (a *Adapter) ConnectBleSimple(session Session, addr Address) (BleConnection, error) {
	return a.ConnectBle(session, addr, ParameterBalanced, RoleClient, PriorityMedium)
}

func (a *Adapter) DisconnectBle(conn BleConnection) error {
	status := C.bleDisconnect(*conn.ptr)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Failed to disconnect from BLE device, error %d", status)
	}

	return nil
}

func (a *Adapter) DiscoverServices(session Session, conn BleConnection) error {
	status := C.bleDiscoverAllServices(*session.ptr, *conn.ptr)
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Failed to discover services from BLE device, error %d", status)
	}

	return nil
}

func (a *Adapter) RetrieveGattDatabase(conn BleConnection) (GattService, error) {
	var outPtr *C.bleGattsService_t
	var serviceCount C.uint32_t

	status := C.bleGetDatabase(*conn.ptr, &outPtr, &serviceCount)
	if status != C.ACE_STATUS_OK {
		return GattService{}, fmt.Errorf(
			"Failed to retrieve GATT database from BLE device, error %d", status,
		)
	}

	return GattService{ptr: outPtr, noSvc: uint(serviceCount)}, nil
}

func (a *Adapter) CleanupGattDatabase(service GattService) error {
	status := C.bleCleanupGattService(service.ptr, C.int(service.noSvc))
	if status != C.ACE_STATUS_OK {
		return fmt.Errorf("Error cleaning up GATT DB: %d", status)
	}
	return nil
}

func (a *Adapter) ReadCharacteristic(
	session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
) (CharacteristicValue, error) {
	rec := C.utilsFindCharRec(
		db.ptr, C.uint32_t(db.noSvc), uuid.cUuid(), C.uint8_t(len(uuid.Bytes)),
	)
	if rec == nil {
		return CharacteristicValue{}, fmt.Errorf("Characteristic %v not found", uuid)
	}

	ch := make(chan readCharacteristicCallback, 1)
	pendingReads.Store(*conn.ptr, ch)

	status := C.bleReadCharacteristic(*session.ptr, *conn.ptr, rec.value)
	if status != C.ACE_STATUS_OK {
		pendingReads.Delete(*conn.ptr)
		return CharacteristicValue{}, fmt.Errorf(
			"Couldn't read Characteristic %v, error: %d", uuid, status,
		)
	}

	select {
	case res := <-ch:
		if res.status != C.ACE_STATUS_OK {
			return CharacteristicValue{}, fmt.Errorf(
				"Read Characteristic %v failed, error: %d", uuid, res.status,
			)
		}

		return newCharacteristicValueFromC(res.value), nil

	case <-time.After(callbackTimeout):
		return CharacteristicValue{}, fmt.Errorf(
			"Timed out reading Characteristic %v", uuid,
		)
	}
}

func (a *Adapter) WriteCharacteristicWithoutResponse(
	session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
	value CharacteristicValueKind,
) error {
	_, err := a.writeCharacteristic(
		session, conn, db, uuid, value, C.ACEBT_BLE_WRITE_TYPE_RESP_NO,
	)
	return err
}

func (a *Adapter) WriteCharacteristicWithResponse(
	session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
	value CharacteristicValueKind,
) (CharacteristicValue, error) {
	return a.writeCharacteristic(
		session, conn, db, uuid, value, C.ACEBT_BLE_WRITE_TYPE_RESP_REQUIRED,
	)
}

func (a *Adapter) writeCharacteristic(
	session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
	value CharacteristicValueKind, responseRequired C.responseType_t,
) (CharacteristicValue, error) {
	rec := C.utilsFindCharRec(
		db.ptr, C.uint32_t(db.noSvc), uuid.cUuid(), C.uint8_t(len(uuid.Bytes)),
	)
	if rec == nil {
		return CharacteristicValue{}, fmt.Errorf("Characteristic %v not found", uuid)
	}

	ch := make(chan writeCharacteristicCallback, 1)
	pendingWrites.Store(*conn.ptr, ch)

	blobUsed := false

	switch v := value.(type) {
	case CharacteristicValueBlob:
		// If blob, reset the shared blob before writes
		C.freeGattBlob(&rec.value)
		data := C.CBytes(v.V)
		defer C.free(data)
		C.setGattBlobFromBytes(&rec.value, (*C.uint8_t)(data), C.uint16_t(len(v.V)))
		rec.value.format = C.BLE_FORMAT_BLOB
		blobUsed = true
	case CharacteristicValueUint8:
		*(*C.uint8_t)(unsafe.Pointer(&rec.value)) = C.uint8_t(v.V)
		rec.value.format = C.BLE_FORMAT_UINT8
	}

	if blobUsed == true {
		defer C.freeGattBlob(&rec.value)
	}

	status := C.bleWriteCharacteristic(*session.ptr, *conn.ptr, &rec.value, responseRequired)
	if status != C.ACE_STATUS_OK {
		pendingReads.Delete(*conn.ptr)
		return CharacteristicValue{}, fmt.Errorf(
			"Couldn't write Characteristic %v, error: %d", uuid, status,
		)
	}

	select {
	case res := <-ch:
		if res.status != C.ACE_STATUS_OK {
			return CharacteristicValue{}, fmt.Errorf(
				"Write Characteristic %v failed, error: %d", uuid, res.status,
			)
		}

		if responseRequired == C.ACEBT_BLE_WRITE_TYPE_RESP_NO {
			return CharacteristicValue{}, nil
		}

		return newCharacteristicValueFromC(res.value), nil

	case <-time.After(callbackTimeout):
		return CharacteristicValue{}, fmt.Errorf(
			"Timed out writing Characteristic %v", uuid,
		)
	}
}

func (a *Adapter) NotifyCharacteristic(
	session Session, conn BleConnection, db GattService, uuid CharacteristicUuid,
) (*Notification, error) {
	rec := C.utilsFindCharRec(
		db.ptr, C.uint32_t(db.noSvc), uuid.cUuid(), C.uint8_t(len(uuid.Bytes)),
	)
	if rec == nil {
		return &Notification{}, fmt.Errorf("Characteristic %v not found", uuid)
	}

	callbackCh := make(chan notifyCharacteristicCallback, 20)
	pendingNotifications.Store(*conn.ptr, callbackCh)

	generatorCh := make(chan CharacteristicValue)
	generatorDone := make(chan struct{})

	go func() {
		defer close(generatorCh)
		for {
			select {
			case <-generatorDone:
				return
			case res := <-callbackCh:
				val := newCharacteristicValueFromC(res.value)
				generatorCh <- val
			}
		}
	}()

	status := C.bleSetNotification(*session.ptr, *conn.ptr, rec.value, C.bool(true))
	if status != C.ACE_STATUS_OK {
		pendingNotifications.Delete(*conn.ptr)
		close(generatorDone)
		return &Notification{}, fmt.Errorf(
			"Couldn't enable notifications on Characteristic %v", uuid,
		)
	}

	stopFn := func() error {
		status := C.bleSetNotification(*session.ptr, *conn.ptr, rec.value, C.bool(false))
		if status != C.ACE_STATUS_OK {
			return fmt.Errorf(
				"Couldn't disable notifications on Characteristic %v", uuid,
			)
		}
		pendingNotifications.Delete(*conn.ptr)
		close(generatorDone)
		return nil
	}

	return &Notification{C: generatorCh, Stop: stopFn}, nil
}
