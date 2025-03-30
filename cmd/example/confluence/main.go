package main

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/example"

func main() {
	if true {
		example.Page()
	}

	if false {
		example.Space()
		example.Search()
		example.User()
		example.Favorite()
		example.Label()
	}
}
