package main

import "github.com/funtimecoding/go-library/pkg/tool/gohab"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gohab.Main(Version, GitHash, BuildDate)
}
