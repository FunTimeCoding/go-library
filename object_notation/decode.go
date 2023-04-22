package object_notation

import "encoding/json"

func Decode(
	value string,
	structure any,
) error {
	return json.Unmarshal([]byte(value), structure)
}
