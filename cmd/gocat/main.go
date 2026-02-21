package main

import "github.com/funtimecoding/go-library/pkg/tool/gocat"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gocat.Main(Version, GitHash, BuildDate)
}
