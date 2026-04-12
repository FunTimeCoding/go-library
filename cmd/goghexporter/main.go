package main

import "github.com/funtimecoding/go-library/pkg/tool/goghexporter"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goghexporter.Main(Version, GitHash, BuildDate)
}
