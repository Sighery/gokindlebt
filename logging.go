package gokindlebt

/*
#include <kindlebt/kindlebt_log.h>
*/
import "C"

import (
	"log/slog"
)

func SetLogLevel(level slog.Level) {
	slog.SetLogLoggerLevel(level)

	switch level {
	case slog.LevelDebug:
		C.kindlebt_set_log_level(C.LOG_LEVEL_TRACE)
	case slog.LevelInfo:
		C.kindlebt_set_log_level(C.LOG_LEVEL_INFO)
	case slog.LevelWarn:
		C.kindlebt_set_log_level(C.LOG_LEVEL_WARN)
	case slog.LevelError:
		C.kindlebt_set_log_level(C.LOG_LEVEL_ERROR)
	}
}
