package system

import (
	"bufio"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func ReadLine(r *bufio.Reader) string {
	result, e := r.ReadString('\n')
	errors.PanicOnError(e)

	return result
}
