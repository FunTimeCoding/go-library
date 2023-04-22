package object_notation

import "encoding/json"

func DecodeSafe(
	value string,
	structure any,
) {
	e := json.Unmarshal([]byte(value), structure)

	if e != nil {
		// do nothing
	}
}
