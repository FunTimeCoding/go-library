package main

import "github.com/funtimecoding/go-library/pkg/tool/gobuild"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gobuild.Main(Version, GitHash, BuildDate)
}
