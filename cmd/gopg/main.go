package main

import "github.com/funtimecoding/go-library/pkg/tool/gopg"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopg.Main(Version, GitHash, BuildDate)
}
