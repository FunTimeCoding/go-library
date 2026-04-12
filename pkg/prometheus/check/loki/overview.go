package loki

import "time"

type overview struct {
	Namespace string
	Count     int
	Latest    time.Time
}
