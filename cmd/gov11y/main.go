package main

import "github.com/funtimecoding/go-library/pkg/tool/gov11y"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gov11y.Main(Version, GitHash, BuildDate)
}
