#ifndef GOKINDLEBT_UTILS_H
#define GOKINDLEBT_UTILS_H

#include <kindlebt/kindlebt.h>

struct aceBT_gattCharRec_t* findCharacteristic(
	bleGattsService_t* services, uint32_t noSvcs, uuid_t uuid, uint8_t uuid_len
);

#endif // GOKINDLEBT_UTILS_H
