package main

import "github.com/funtimecoding/go-library/pkg/tool/gogitlab"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogitlab.Main(Version, GitHash, BuildDate)
}
