package main

import "github.com/funtimecoding/go-library/pkg/tool/gobrew"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gobrew.Main(Version, GitHash, BuildDate)
}
