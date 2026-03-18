package main

import "github.com/funtimecoding/go-library/pkg/tool/goc"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goc.Main(Version, GitHash, BuildDate)
}
