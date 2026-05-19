package model_context

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (s *Server) listMemoriesWithContent(tag string) ([]store.Memory, error) {
	summaries, e := s.service.Store.ListMemories("", tag, true)

	if e != nil {
		return nil, e
	}

	var result []store.Memory

	for _, sum := range summaries {
		m, e := s.service.Store.GetMemory(sum.Identifier)

		if e != nil {
			continue
		}

		result = append(result, *m)
	}

	return result, nil
}
