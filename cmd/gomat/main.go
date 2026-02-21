package main

import "github.com/funtimecoding/go-library/pkg/tool/gomat"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomat.Main(Version, GitHash, BuildDate)
}
