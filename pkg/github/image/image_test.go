package image

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v84/github"
	"testing"
	"time"
)

func TestImage(t *testing.T) {
	meta, _ := json.Marshal(
		github.PackageMetadata{
			PackageType: new(constant.ContainerPackageType),
			Container: &github.PackageContainerMetadata{
				Tags: []string{strings.Alfa},
			},
		},
	)
	i := New(
		&github.PackageVersion{
			ID:        new(int64(1)),
			Name:      new(strings.Bravo),
			CreatedAt: &github.Timestamp{},
			Metadata:  meta,
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&Image{
			Identifier: 1,
			Digest:     strings.Bravo,
			Tags:       []string{strings.Alfa},
			Create:     time.Time{},
		},
		i,
	)
}
