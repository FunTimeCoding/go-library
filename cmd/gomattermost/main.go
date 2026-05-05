package main

import "github.com/funtimecoding/go-library/pkg/tool/gomattermost"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomattermost.Main(Version, GitHash, BuildDate)
}
