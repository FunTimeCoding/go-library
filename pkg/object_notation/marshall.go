package object_notation

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Marshall(a any) []byte {
	result, e := json.Marshal(a)
	errors.PanicOnError(e)

	return result
}
