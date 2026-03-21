package main

import (
	brave "github.com/funtimecoding/go-library/pkg/brave/example"
	chromium "github.com/funtimecoding/go-library/pkg/chromium/example"
)

func main() {
	chromium.Tab()

	if false {
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
