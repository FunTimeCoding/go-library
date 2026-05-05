// Start: cd py/iterm_bridge && uv run main.py
package main

import "github.com/funtimecoding/go-library/pkg/tool/goitermd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goitermd.Main(Version, GitHash, BuildDate)
}
