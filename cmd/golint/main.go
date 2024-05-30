package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/lint"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	for _, element := range system.EmptyDirectories(
		constant.CurrentDirectory,
	) {
		fmt.Printf("Empty directory: %s\n", element)
	}

	for _, element := range system.GoFiles(
		constant.CurrentDirectory,
	) {
		f := system.Open(element)

		if linted := lint.Fix(f); linted != "" {
			fmt.Printf("Simplify import %s\n", element)
			system.SaveFile(element, linted)
		}

		errors.LogClose(f)
	}
}
