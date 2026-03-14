package notation

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func MarshalIndentBytes(a any) []byte {
	result, e := json.MarshalIndent(a, "", "\t")
	errors.PanicOnError(e)

	return result
}
