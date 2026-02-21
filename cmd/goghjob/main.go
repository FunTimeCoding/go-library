package main

import "github.com/funtimecoding/go-library/pkg/tool/goghjob"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goghjob.Main(Version, GitHash, BuildDate)
}
