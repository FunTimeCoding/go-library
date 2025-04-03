package main

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/example"

func main() {
	if true {
		// TODO: Load all pages in a space and save as JSON, with markdown text
		// TODO: Load all sub-pages of a page and save as JSON, with markdown text
		example.Page()
	}

	if false {
		example.Label()
		example.User()
		example.Favorite()
		example.Search()
		example.Space()
	}
}
