package response

import "time"

type Release struct {
	Version      string     `json:"version"`
	ShortVersion string     `json:"shortVersion"`
	DateCreated  *time.Time `json:"dateCreated"`
	DateReleased *time.Time `json:"dateReleased"`
	FirstEvent   *time.Time `json:"firstEvent"`
	LastEvent    *time.Time `json:"lastEvent"`
	NewGroups    int        `json:"newGroups"`
}
