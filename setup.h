#ifndef GOKINDLEBT_SETUP_H
#define GOKINDLEBT_SETUP_H

#include <kindlebt/kindlebt.h>

sessionHandle newSessionHandle(void);
bleConnHandle newBleConnHandle(void);
bleGattsService_t* newBleGattsService(void);

#endif // GOKINDLEBT_SETUP_H
