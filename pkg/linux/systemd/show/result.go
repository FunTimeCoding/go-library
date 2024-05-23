package show

import "time"

type Result struct {
	ActiveState string
	SubState    string
	ActiveEnter time.Time
	// ExecMainStart May not be set for every service
	ExecMainStart time.Time
}
