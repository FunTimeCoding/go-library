package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	fmt.Printf("Executable: %s\n", system.ExecutablePath())
}
