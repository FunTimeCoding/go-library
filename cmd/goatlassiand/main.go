package main

import "github.com/funtimecoding/go-library/pkg/tool/goatlassiand"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goatlassiand.Main(Version, GitHash, BuildDate)
}
