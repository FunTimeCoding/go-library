package main

import "github.com/funtimecoding/go-library/pkg/tool/gosed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosed.Main(Version, GitHash, BuildDate)
}
