package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	TokenEnvironment = "GITHUB_TOKEN"
	RunEnvironment   = "GITHUB_RUN_ID"
	ReferenceName    = "GITHUB_REF_NAME"

	DelveNamespace  = "go-delve"
	DelveRepository = "delve"

	LibraryNamespace  = "funtimecoding"
	LibraryRepository = "go-library"

	SignatureHeader = "X-Hub-Signature-256"
)

var Format = option.ExtendedColor.Copy()
