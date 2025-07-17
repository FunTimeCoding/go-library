package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/text/option"
)

const (
	Go      = "go"
	Mod     = "mod"
	Tidy    = "tidy"
	Edit    = "edit"
	Build   = "build"
	Get     = "get"
	Version = "version"

	LinkerFlagsArgument = "-ldflags"
	OutputArgument      = "-o"
	TagsArgument        = "-tags"

	VersionArgument = "-go"

	LinkerSetVariable = "-X"

	NativeEnabled = "CGO_ENABLED"
	System        = "GOOS"
	Architecture  = "GOARCH"
	Proxy         = "GOPROXY"

	False  = "0"
	Direct = "direct"

	CurrentDirectory  = "."
	GoExtension       = ".go"
	MarkupExtension   = ".yaml"
	MarkdownExtension = ".md"
)

var (
	StartOfTime       = assert.NewDay(1)
	CompactWhitespace = option.Compact()
)
