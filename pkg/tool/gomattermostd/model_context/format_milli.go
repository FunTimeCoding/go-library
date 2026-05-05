package model_context

import "time"

func formatMilli(v int64) string {
	return formatTime(time.UnixMilli(v))
}
