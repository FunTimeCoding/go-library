package main

import "github.com/funtimecoding/go-library/pkg/tool/gobundle"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gobundle.Main(Version, GitHash, BuildDate)
}
