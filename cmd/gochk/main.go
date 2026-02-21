package main

import "github.com/funtimecoding/go-library/pkg/tool/gochk"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gochk.Main(Version, GitHash, BuildDate)
}
