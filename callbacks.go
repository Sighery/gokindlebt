package gokindlebt

/*
#include <kindlebt/kindlebt.h>
#include "callbacks.h"
*/
import "C"

import (
	"log/slog"
)

type readCharacteristicCallback struct {
	conn   C.bleConnHandle
	value  C.bleGattCharacteristicsValue_t
	status C.status_t
}

//export goOnBleGattcReadCharsCallback
func goOnBleGattcReadCharsCallback(
	connHandle C.bleConnHandle, charValue C.bleGattCharacteristicsValue_t, status C.status_t,
) {
	chAny, loaded := pendingReads.LoadAndDelete(connHandle)

	if loaded == false {
		slog.Warn("No reads channel registered, this shouldn't happen")
		return
	}

	ch := chAny.(chan readCharacteristicCallback)
	ch <- readCharacteristicCallback{conn: connHandle, value: charValue, status: status}
}

type writeCharacteristicCallback struct {
	conn   C.bleConnHandle
	value  C.bleGattCharacteristicsValue_t
	status C.status_t
}

//export goOnBleGattcWriteCharsCallback
func goOnBleGattcWriteCharsCallback(
	connHandle C.bleConnHandle, charValue C.bleGattCharacteristicsValue_t, status C.status_t,
) {
	chAny, loaded := pendingWrites.LoadAndDelete(connHandle)

	if loaded == false {
		slog.Warn("No writes channel registered, this shouldn't happen")
		return
	}

	ch := chAny.(chan writeCharacteristicCallback)
	ch <- writeCharacteristicCallback{conn: connHandle, value: charValue, status: status}
}

type notifyCharacteristicCallback struct {
	conn  C.bleConnHandle
	value C.bleGattCharacteristicsValue_t
}

//export goOnBleGattcNotifyCharsCallback
func goOnBleGattcNotifyCharsCallback(
	connHandle C.bleConnHandle, charValue C.bleGattCharacteristicsValue_t,
) {
	chAny, loaded := pendingNotifications.Load(connHandle)

	if loaded == false {
		slog.Warn("No notifications channel registered, this shouldn't happen")
		return
	}

	ch := chAny.(chan notifyCharacteristicCallback)
	select {
	case ch <- notifyCharacteristicCallback{conn: connHandle, value: charValue}:
	default:
		slog.Warn("Dropping notification due to full channel")
	}
}
