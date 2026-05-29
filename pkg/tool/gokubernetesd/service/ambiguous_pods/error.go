package ambiguous_pods

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (e *AmbiguousPods) Error() string {
	return fmt.Sprintf(
		"multiple pods match %s: %s",
		e.Name,
		join.CommaSpace(e.Pods),
	)
}
