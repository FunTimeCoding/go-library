package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2"
)

func World() {
	fmt.Printf("Worlds: %+v\n", gw2.NewEnvironment().Worlds())
}
