//go:build local

package main

import "github.com/funtimecoding/go-library/pkg/generative/chroma/example"

func main() {
	if true {
		// TODO: Load Markdown files
		// TODO: Load Mattermost threads
		// TODO: Load Confluence pages
		example.Collection()
	}

	if false {
		example.Client()
	}
}
