package main

import "github.com/funtimecoding/go-library/pkg/tool/goclean"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goclean.Main(Version, GitHash, BuildDate)
}
