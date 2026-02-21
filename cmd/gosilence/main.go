package main

import "github.com/funtimecoding/go-library/pkg/tool/gosilence"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosilence.Main(Version, GitHash, BuildDate)
}
