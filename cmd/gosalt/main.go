package main

import "github.com/funtimecoding/go-library/pkg/tool/gosalt"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosalt.Main(Version, GitHash, BuildDate)
}
