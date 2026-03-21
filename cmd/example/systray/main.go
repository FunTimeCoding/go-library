//go:build local

package main

import (
	"fmt"
	"github.com/getlantern/systray"
)

func main() {
	fmt.Println("Start")
	systray.Run(onReady, onExit)
	fmt.Println("After run")
}
