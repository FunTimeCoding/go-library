package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/markdown/file"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func main() {
	// TODO: Design an object to hold a markdown runbook
	base := environment.Get("WIKI_PATH")

	for _, n := range system.Files(base) {
		fmt.Printf("File: %s\n", n)
		source := system.ReadBytes(system.Join(base, n))
		fmt.Println(
			console.Cyan(
				"%s",
				strings.PrefixMultiline(string(source), "> "),
			),
		)
		f := file.New(&source)
		f.Parse()

		if true {
			break
		}
	}
}
