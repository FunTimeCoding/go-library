package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/sound"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Sound() {
	fmt.Println(system.Run(sound.Afplay, sound.SosumiPath))
	fmt.Println(system.Run(sound.Afplay, sound.TinkPath))
}
