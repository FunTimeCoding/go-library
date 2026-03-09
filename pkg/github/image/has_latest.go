package image

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"slices"
)

func HasLatest(v *Image) bool {
	return slices.Contains(v.Tags, constant.LatestVersion)
}
