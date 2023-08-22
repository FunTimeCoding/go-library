package object_notation

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func MarshallIndent(a any) []byte {
	result, e := json.MarshalIndent(a, "", "\t")
	errors.PanicOnError(e)

	return result
}
