package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/telegram"
	"github.com/funtimecoding/go-library/pkg/telegram/constant"
)

func Update() {
	c := telegram.NewEnvironment()
	f := constant.Format

	for _, m := range c.Messages() {
		fmt.Println(m.Format(f))
	}
}
