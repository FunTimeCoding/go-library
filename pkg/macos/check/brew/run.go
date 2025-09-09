package brew

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/macos/brew"
	"github.com/funtimecoding/go-library/pkg/macos/check/brew/option"
)

func Run(o *option.Brew) {
	p := brew.New().Outdated()

	if o.Notation {
		printNotation(o)
	}

	fmt.Printf("Outdated: %+v\n", p)
}
