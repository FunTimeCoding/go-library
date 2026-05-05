package main

import "github.com/funtimecoding/go-library/pkg/tool/gopostgresd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopostgresd.Main(Version, GitHash, BuildDate)
}
