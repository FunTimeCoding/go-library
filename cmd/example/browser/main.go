package main

import (
	brave "github.com/funtimecoding/go-library/pkg/brave/example"
	chromium "github.com/funtimecoding/go-library/pkg/chromium/example"
	anthropic "github.com/funtimecoding/go-library/pkg/generative/anthropic/site/example"
)

func main() {
	anthropic.Dump()

	if false {
		chromium.Tab()
		brave.BookmarkSearch()
		brave.BookmarkNode()
		brave.BookmarkFile()
		brave.Extract()
		brave.Open()
		brave.Send()
		brave.Profile()
		chromium.Tabs()
		chromium.Tab()
		chromium.Open()
	}
}
