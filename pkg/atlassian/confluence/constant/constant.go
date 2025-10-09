package constant

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

const (
	DefaultSpaceEnvironment = "CONFLUENCE_DEFAULT_SPACE"
	DefaultPageEnvironment  = "CONFLUENCE_DEFAULT_PAGE"
	LabelEnvironment        = "CONFLUENCE_LABEL"

	NoSpace = "no space"

	PageType = "page"

	PathPrefix = "/wiki/api/v2"
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
	Format = option.ExtendedColor.Copy().Tag(tag.Link)
	Dense  = option.Color.Copy().Tag(tag.Link, tag.Dense)
)
