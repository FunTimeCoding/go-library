package main

import "github.com/funtimecoding/go-library/pkg/tool/goclauded"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goclauded.Main(Version, GitHash, BuildDate)
}
