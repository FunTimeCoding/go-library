package main

import "github.com/funtimecoding/go-library/pkg/tool/gosproutd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosproutd.Main(Version, GitHash, BuildDate)
}
