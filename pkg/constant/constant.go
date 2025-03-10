package constant

import "time"

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

	CurrentDirectory = "."
	GoExtension      = ".go"
	MarkupExtension  = ".yaml"
)

var FixtureDate = time.Date(
	2020,
	1,
	1,
	0,
	0,
	0,
	0,
	time.UTC,
)
