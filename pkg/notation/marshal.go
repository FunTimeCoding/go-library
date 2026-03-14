package notation

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Marshal(a any) []byte {
	result, e := json.Marshal(a)
	errors.PanicOnError(e)

	return result
}
