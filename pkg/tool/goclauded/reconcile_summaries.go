package goclauded

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
)

func reconcileSummaries(
	s *store.Store,
	i face.Indexer,
) {
	for _, m := range s.ListSummaries() {
		i.MustPush(m.Name, m.Body)
	}
}
