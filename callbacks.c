#include "callbacks.h"

#include <kindlebt/kindlebt.h>

__attribute__((unused)) void setGcCallbacks(aceBT_bleGattClientCallbacks_t* cb) {
	cb->on_ble_gattc_read_characteristics_cb = goOnBleGattcReadCharsCallback;
	cb->on_ble_gattc_write_characteristics_cb = goOnBleGattcWriteCharsCallback;
	cb->notify_characteristics_cb = goOnBleGattcNotifyCharsCallback;
}
