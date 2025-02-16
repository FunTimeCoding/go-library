package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func Project() {
	for _, p := range sentry.NewEnvironment().Projects() {
		fmt.Printf("%+v\n", p)
	}
}
