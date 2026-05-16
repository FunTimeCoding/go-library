package firefox

import "encoding/json"

func decodeResult(
	r *reply,
	v any,
) error {
	return json.Unmarshal(r.Result, v)
}
