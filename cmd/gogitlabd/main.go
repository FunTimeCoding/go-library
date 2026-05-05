package main

import "github.com/funtimecoding/go-library/pkg/tool/gogitlabd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogitlabd.Main(Version, GitHash, BuildDate)
}
