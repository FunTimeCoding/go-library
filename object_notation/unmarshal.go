package object_notation

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/errors"
)

func Unmarshal(
	value string,
	structure any,
) {
	errors.PanicOnError(json.Unmarshal([]byte(value), structure))
}
