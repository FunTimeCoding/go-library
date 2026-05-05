package main

import "github.com/funtimecoding/go-library/pkg/tool/gosentryd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosentryd.Main(Version, GitHash, BuildDate)
}
