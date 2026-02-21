package main

import "github.com/funtimecoding/go-library/pkg/tool/gokevt"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gokevt.Main(Version, GitHash, BuildDate)
}
