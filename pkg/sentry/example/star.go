package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func Star() {
	// TODO: API does not expose stars. Read from the database?
	fmt.Println(sentry.NewEnvironment().Favourites())
}
