package message

import "encoding/json"

func parseNotation(s string) (map[string]any, bool) {
	var result map[string]any

	if e := json.Unmarshal([]byte(s), &result); e != nil {
		return nil, false
	}

	return result, true
}
