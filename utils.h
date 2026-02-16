#ifndef GOKINDLEBT_UTILS_H
#define GOKINDLEBT_UTILS_H

#include <kindlebt/kindlebt.h>

struct aceBT_gattCharRec_t* findCharacteristic(
	bleGattsService_t* services, uint32_t noSvcs, uuid_t uuid, uint8_t uuid_len
);

// Here because CGO can't handle C functions with variadic arguments
char* append_to_buffer_wrapper(
	char* buf, size_t* size, size_t* offset, const char* fmt, uint32_t index, void* ptr
);

#endif // GOKINDLEBT_UTILS_H
