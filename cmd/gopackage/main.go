package main

import "github.com/funtimecoding/go-library/pkg/tool/gopackage"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopackage.Main(Version, GitHash, BuildDate)
}
