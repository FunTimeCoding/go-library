package main

import "github.com/funtimecoding/go-library/pkg/tool/gonetbd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gonetbd.Main(Version, GitHash, BuildDate)
}
