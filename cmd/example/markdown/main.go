package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/markdown/file"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
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

		p := goldmark.DefaultParser()
		o := p.Parse(text.NewReader(source))

		if false {
			file.Walk(&source, o)
		}

		if true {
			file.WalkTree(&source, o)
		}

		if true {
			break
		}
	}
}
