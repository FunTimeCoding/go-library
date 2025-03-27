package constant

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

const NoSpace = "no space"

var DenseFormat = option.Color.Copy().Tag(tag.Link, tag.Dense)
