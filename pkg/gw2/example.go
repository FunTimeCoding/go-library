package gw2

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Example() {
	c := New(environment.Get(TokenEnvironment))
	fmt.Printf("Worlds: %+v\n", c.Worlds())
	fmt.Printf("Account: %+v\n", c.Account())
}
