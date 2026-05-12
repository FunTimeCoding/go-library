package main

import "github.com/funtimecoding/go-library/pkg/tool/gotelemetryd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gotelemetryd.Main(Version, GitHash, BuildDate)
}
