package mock_indexer

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/face/search_option"
)

func (i *Indexer) Search(_ *search_option.Option) ([]face.SearchResult, error) {
	return nil, nil
}
