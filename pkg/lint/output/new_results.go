package output

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
)

func NewResults() *Results {
	return &Results{
		workDirectory: fmt.Sprintf(
			"%s/",
			system.WorkDirectory(),
		),
	}
}
