package main

import "github.com/funtimecoding/go-library/pkg/tool/goclaude"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goclaude.Main(Version, GitHash, BuildDate)
}
