package gokindlebt

/*
#include <kindlebt/kindlebt.h>
*/
import "C"

import (
	"sync"
	"time"
)

const (
	callbackTimeout = 10 * time.Second
)

var (
	// Need this to convert to BleConnection later on in callbacks
	connRegistry = make(map[C.bleConnHandle]*BleConnection)

	// Registry for callback operations
	pendingReads         sync.Map
	pendingWrites        sync.Map
	pendingNotifications sync.Map
	pendingDbs           sync.Map
)
