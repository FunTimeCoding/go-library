package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	TokenEnvironment = "GITHUB_TOKEN"
	RunEnvironment   = "GITHUB_RUN_ID"

	DelveNamespace  = "go-delve"
	DelveRepository = "delve"

	LibraryNamespace  = "funtimecoding"
	LibraryRepository = "go-library"
)

var Format = option.ExtendedColor.Copy()
