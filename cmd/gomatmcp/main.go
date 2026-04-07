package main

import "github.com/funtimecoding/go-library/pkg/tool/gomatmcp"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomatmcp.Main(Version, GitHash, BuildDate)
}
