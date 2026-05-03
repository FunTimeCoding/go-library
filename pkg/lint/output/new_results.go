package output

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
)

func NewResults() Results {
	return Results{
		workingDirectory: fmt.Sprintf(
			"%s/",
			system.WorkingDirectory(),
		),
	}
}
