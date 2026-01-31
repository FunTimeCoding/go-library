package brew

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system/macos/brew"
	"github.com/funtimecoding/go-library/pkg/system/macos/check/brew/option"
)

func Check(o *option.Brew) {
	b := brew.New()

	if o.Notation {
		printNotation(o)
	}

	if true {
		// TODO: Unmarshal fails
		fmt.Printf("Installed: %+v\n", b.Installed())
	}

	if false {
		fmt.Printf("Outdated: %+v\n", b.Outdated())
	}
}
