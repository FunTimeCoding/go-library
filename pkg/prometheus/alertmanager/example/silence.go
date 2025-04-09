package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func Silence() {
	f := option.Color.Copy().Tag(tag.State)

	for _, a := range internal.Alertmanager().Silences(true) {
		fmt.Println(a.Format(f))
	}
}
