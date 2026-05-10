package response

import "encoding/json"

type MinionList struct {
	Return []map[string]json.RawMessage `json:"return"`
}
