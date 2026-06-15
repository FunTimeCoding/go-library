package memory_indexer_tester

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func EmptyFixture() http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		_ *http.Request,
	) {
		web.EncodeNotation(
			w,
			client.SearchOutcome{Results: []client.SearchResult{}},
		)
	}
}
