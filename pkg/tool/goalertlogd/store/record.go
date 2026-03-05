package store

import "time"

type Record struct {
	Fingerprint string            `json:"fingerprint"`
	Name        string            `json:"name"`
	Severity    string            `json:"severity"`
	Summary     string            `json:"summary"`
	Labels      map[string]string `json:"labels"`
	Start       time.Time         `json:"start"`
	End         *time.Time        `json:"end,omitempty"`
}
