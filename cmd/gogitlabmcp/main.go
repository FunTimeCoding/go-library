package main

import "github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogitlabmcp.Main(Version, GitHash, BuildDate)
}
