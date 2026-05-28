package process

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"strings"
)

func (p *Process) Output() string {
	if p.stdout == nil {
		return ""
	}

	b, e := io.ReadAll(p.stdout)
	errors.PanicOnError(e)

	return strings.TrimSpace(string(b))
}
