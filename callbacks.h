#ifndef GOKINDLEBT_CALLBACKS_H
#define GOKINDLEBT_CALLBACKS_H

#include <kindlebt/kindlebt.h>

// Declaration of Go callbacks. Will be defined in Go code
extern void goOnBleGattcReadCharsCallback(bleConnHandle, bleGattCharacteristicsValue_t, status_t);
extern void goOnBleGattcWriteCharsCallback(bleConnHandle, bleGattCharacteristicsValue_t, status_t);
extern void goOnBleGattcNotifyCharsCallback(bleConnHandle, bleGattCharacteristicsValue_t);
extern void goOnBleGattcGetDbCallback(bleConnHandle, bleGattsService_t*, uint32_t);

// Callback setters because it seems I cannot just assign from Go
__attribute__((unused)) void setGcCallbacks(bleGattClientCallbacks_t*);

#endif // GOKINDLEBT_CALLBACKS_H
