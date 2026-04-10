// Setup: ln -s $HOME/src/go-library/py/sublime_bridge "$HOME/Library/Application Support/Sublime Text/Packages/SublimeBridge"
package main

import "github.com/funtimecoding/go-library/pkg/tool/gosublmcp"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosublmcp.Main(Version, GitHash, BuildDate)
}
