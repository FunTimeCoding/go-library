package main

import "github.com/funtimecoding/go-library/pkg/tool/gosourced"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosourced.Main(Version, GitHash, BuildDate)
}
