package main

import "github.com/funtimecoding/go-library/pkg/tool/goatld"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goatld.Main(Version, GitHash, BuildDate)
}
