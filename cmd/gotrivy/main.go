package main

import "github.com/funtimecoding/go-library/pkg/tool/gotrivy"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gotrivy.Main(Version, GitHash, BuildDate)
}
