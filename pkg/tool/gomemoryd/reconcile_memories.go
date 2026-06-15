package gomemoryd

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/service"
)

func reconcileMemories(s *service.Service) {
	memories, e := s.ListMemories("", "", true)
	errors.PanicOnError(e)

	for _, m := range memories {
		full, f := s.GetMemory(m.Identifier)
		errors.PanicOnError(f)
		s.MustIndexMemory(
			fmt.Sprintf("memory/%d", full.Identifier),
			full.Content,
			map[string]string{
				"memory_id":   fmt.Sprintf("%d", full.Identifier),
				"type":        full.Type,
				"name":        full.Name,
				"description": full.Description,
			},
		)
	}
}
