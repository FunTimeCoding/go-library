package tag

import "time"

type response struct {
	Name        string     `json:"name"`
	LastUpdated *time.Time `json:"last_updated"`
}
