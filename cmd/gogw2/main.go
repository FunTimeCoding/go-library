package main

import "github.com/funtimecoding/go-library/pkg/tool/gogw2"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogw2.Main(Version, GitHash, BuildDate)
}
