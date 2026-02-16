#include "utils.h"

#include <stdio.h>

#include <kindlebt/kindlebt.h>
#include <kindlebt/kindlebt_log.h>

struct aceBT_gattCharRec_t* findCharacteristic(
	bleGattsService_t* services, uint32_t noSvcs, uuid_t uuid, uint8_t uuid_len
) {
	struct aceBT_gattCharRec_t* char_rec = NULL;

	if (!services) {
		perror("Gatt DB has not been populated yet");
		return (NULL);
	}

	for (uint32_t i = 0; i < noSvcs; i++) {
		bleGattsService_t* service = &services[i];

		STAILQ_FOREACH(char_rec, &service->charsList, link) {
			if (!memcmp(char_rec->value.gattRecord.uuid.uu, &uuid.uu, uuid_len)) {
				return (char_rec);
			}
		}
	}

	perror("GATT Characteristic UUID could not be found!");
	return (NULL);
}

char* append_to_buffer_wrapper(
	char* buf, size_t* size, size_t* offset, const char* fmt, uint32_t index, void* ptr
) {
	return append_to_buffer(buf, size, offset, fmt, index, ptr);
}
