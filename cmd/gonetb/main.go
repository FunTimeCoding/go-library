package main

import "github.com/funtimecoding/go-library/pkg/tool/gonetb"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gonetb.Main(Version, GitHash, BuildDate)
}
