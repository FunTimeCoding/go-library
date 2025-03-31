package constant

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

const NoSpace = "no space"

// Fixture
const (
	OperationsSpace = "Operations"
	ExamplePage     = "Example"
)

// Body format
const (
	ViewFormat      = "view"
	AtlasFormat     = "atlas_doc_format"
	StorageFormat   = "storage"
	ExportFormat    = "export_view"
	AnonymousFormat = "anonymous_export_view"
	StyledFormat    = "styled_view"
	EditFormat      = "editor"
)

var (
	DefaultFormat = option.ExtendedColor
	DenseFormat   = option.Color.Copy().Tag(tag.Link, tag.Dense)
)
