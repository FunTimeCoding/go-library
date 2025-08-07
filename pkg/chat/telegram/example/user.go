package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chat/telegram"
)

func User() {
	c := telegram.NewEnvironment()
	fmt.Printf("User: %s\n", c.Self().UserName)
}
