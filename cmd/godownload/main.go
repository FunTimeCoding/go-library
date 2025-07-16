package main

import "github.com/funtimecoding/go-library/pkg/tool/godownload"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	godownload.Main(Version, GitHash, BuildDate)
}
