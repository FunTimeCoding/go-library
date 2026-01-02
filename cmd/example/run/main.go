package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	w := system.WorkingDirectory()
	fmt.Printf("Directory: %s\n", w)

	if !runtime.RunningFromBinary() {
		fmt.Println("Run from source")
	} else {
		fmt.Println("Run from binary")
	}
}
