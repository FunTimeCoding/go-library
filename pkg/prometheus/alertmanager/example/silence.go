package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
)

func Silence() {
	f := format.Color.Copy().Tag(tag.State)

	for _, a := range alertmanager.NewEnvironment().Silences(true) {
		fmt.Println(a.Format(f))
	}
}
