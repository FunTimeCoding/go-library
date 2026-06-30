package face

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/face/search_option"

type Searcher interface {
	Search(
		o *search_option.Option,
	) ([]SearchResult, error)
}
