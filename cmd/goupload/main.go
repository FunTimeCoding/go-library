package main

import "github.com/funtimecoding/go-library/pkg/tool/goupload"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goupload.Main(Version, GitHash, BuildDate)
}
