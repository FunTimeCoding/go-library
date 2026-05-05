package main

import "github.com/funtimecoding/go-library/pkg/tool/gonetboxd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gonetboxd.Main(Version, GitHash, BuildDate)
}
