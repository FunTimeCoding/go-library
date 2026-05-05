package monitor

import "time"

func nowMilli() int64 {
	return time.Now().UnixMilli()
}
