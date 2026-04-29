package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/gitlab"
)

func Runner() {
	g := gitlab.NewEnvironment()
	f := option.ExtendedColor.Copy().Raw()

	for _, r := range g.MustRunners(true) {
		r = g.MustRunner(r.Identifier)
		fmt.Println(r.Format(f))
	}
}
