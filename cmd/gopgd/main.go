package main

import "github.com/funtimecoding/go-library/pkg/tool/gopgd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopgd.Main(Version, GitHash, BuildDate)
}
