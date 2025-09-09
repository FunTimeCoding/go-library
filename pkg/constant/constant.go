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
	True   = "1"
	Direct = "direct"

	CurrentDirectory     = "."
	GoExtension          = ".go"
	MarkdownExtension    = ".md"
	MarkupExtension      = ".yaml"
	ShortMarkupExtension = ".yml"

	LatestVersion = "latest"

	PhysicalTest0 = "00:00:00:00:00:00"
	PhysicalTest1 = "00:00:00:00:00:01"
	PhysicalTest2 = "00:00:00:00:00:02"
)

// For console status option
const (
	LabelKey = "label"
	TagKey   = "tag"
)

var (
	StartOfTime       = assert.NewDay(1)
	CompactWhitespace = option.Compact()
)
