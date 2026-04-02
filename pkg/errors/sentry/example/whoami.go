package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
)

func Whoami() {
	c := sentry.NewEnvironment()
	u := c.Whoami()
	fmt.Printf("User: %+v\n", u)
}
