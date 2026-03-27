package main

import "github.com/funtimecoding/go-library/pkg/tool/gofix"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gofix.Main(Version, GitHash, BuildDate)
}
