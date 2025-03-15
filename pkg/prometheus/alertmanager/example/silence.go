package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
)

func Silence() {
	for _, a := range alertmanager.NewEnvironment().Silences(true) {
		fmt.Println(a.Format(format.Color))
	}
}
