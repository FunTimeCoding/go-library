package response

import "time"

type TagValue struct {
	Value     string     `json:"value"`
	Count     int        `json:"count"`
	FirstSeen *time.Time `json:"firstSeen"`
	LastSeen  *time.Time `json:"lastSeen"`
}
