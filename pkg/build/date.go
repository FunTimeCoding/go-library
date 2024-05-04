package build

import "time"

func Date() string {
	return time.Now().Format(time.RFC3339)
}
