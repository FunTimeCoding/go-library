package notation

import "encoding/json"

func DecodeBytes(
	value []byte,
	structure any,
) error {
	return json.Unmarshal(value, structure)
}
