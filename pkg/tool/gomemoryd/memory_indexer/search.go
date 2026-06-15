package memory_indexer

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/client"
)

func (i *Indexer) Search(
	query string,
	collection string,
	limit int,
) ([]face.SearchResult, error) {
	full := true
	mode := client.GetSearchParamsMode("hybrid")
	r, e := i.client.GetSearch(
		context.Background(),
		&client.GetSearchParams{
			Query:      query,
			Collection: &collection,
			Limit:      &limit,
			Full:       &full,
			Mode:       &mode,
		},
	)

	if e != nil {
		return nil, e
	}

	parsed, e := client.ParseGetSearchResponse(r)

	if e != nil {
		return nil, e
	}

	if parsed.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %s", parsed.Status())
	}

	var result []face.SearchResult

	for _, s := range parsed.JSON200.Results {
		body := ""

		if s.Body != nil {
			body = *s.Body
		}

		result = append(
			result,
			face.SearchResult{
				Path:  s.Path,
				Body:  body,
				Score: float64(s.Score),
			},
		)
	}

	return result, nil
}
