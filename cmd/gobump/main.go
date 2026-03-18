package main

import "github.com/funtimecoding/go-library/pkg/tool/gobump"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gobump.Main(Version, GitHash, BuildDate)
}
