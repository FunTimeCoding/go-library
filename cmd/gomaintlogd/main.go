package main

import "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomaintlogd.Main(Version, GitHash, BuildDate)
}
