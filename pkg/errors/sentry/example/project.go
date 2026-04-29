package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
)

func Project() {
	for _, p := range sentry.NewEnvironment().MustProjects() {
		fmt.Printf("%+v\n", p)
	}
}
