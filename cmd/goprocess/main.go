package main

import "github.com/funtimecoding/go-library/pkg/tool/goprocess"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goprocess.Main(Version, GitHash, BuildDate)
}
