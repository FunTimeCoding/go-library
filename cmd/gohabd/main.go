package main

import "github.com/funtimecoding/go-library/pkg/tool/gohabd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gohabd.Main(Version, GitHash, BuildDate)
}
