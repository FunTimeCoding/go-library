package maps

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Encode(input map[string]string) []byte {
	result, e := json.Marshal(input)
	errors.PanicOnError(e)

	return result
}
