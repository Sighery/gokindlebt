#include "setup.h"

#include <stdio.h>
#include <stdlib.h>

sessionHandle newSessionHandle(void) {
	sessionHandle s = malloc(sizeof(sessionHandle));
	if (!s) return NULL;
	return s;
}

bleConnHandle newBleConnHandle(void) {
	bleConnHandle c = malloc(sizeof(bleConnHandle));
	if (!c) return NULL;
	return c;
}

bleGattsService_t* newBleGattsService(void) {
	bleGattsService_t* s = malloc(sizeof(bleGattsService_t));
	if (!s) return NULL;
	return s;
}
