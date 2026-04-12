package firefox

import "encoding/json"

func decodeResult(
	r response,
	v any,
) error {
	return json.Unmarshal(r.Result, v)
}
