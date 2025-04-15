package notation

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func DecodeBytes(
	value []byte,
	structure any,
) {
	errors.PanicOnError(json.Unmarshal(value, structure))
}
