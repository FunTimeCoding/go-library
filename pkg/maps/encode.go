package maps

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Encode(m map[string]string) []byte {
	result, e := json.Marshal(m)
	errors.PanicOnError(e)

	return result
}
