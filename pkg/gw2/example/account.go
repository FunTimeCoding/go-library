package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2"
)

func Account() {
	fmt.Printf("Account: %+v\n", gw2.NewEnvironment().Account())
}
