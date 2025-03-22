package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
)

func Silence() {
	f := option.Color.Copy().Tag(tag.State)

	for _, a := range alertmanager.NewEnvironment().Silences(true) {
		fmt.Println(a.Format(f))
	}
}
