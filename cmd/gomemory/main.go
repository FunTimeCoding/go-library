package main

import "github.com/funtimecoding/go-library/pkg/tool/gomemory"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomemory.Main(Version, GitHash, BuildDate)
}
