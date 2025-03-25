package score_colorer

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/math/range_mapping"
)

type Colorer struct {
	assignments map[string]face.SprintFunction
	mapping     []*range_mapping.Mapping
}
