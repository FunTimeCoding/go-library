package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/sound"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	fmt.Println(
		system.Run(sound.Afplay, "/System/Library/Sounds/Tink.aiff"),
	)
}
