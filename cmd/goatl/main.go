package main

import "github.com/funtimecoding/go-library/pkg/tool/goatl"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goatl.Main(Version, GitHash, BuildDate)
}
