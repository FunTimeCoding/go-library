package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (s *Service) CreateMemory(
	name string,
	content string,
	description string,
	memoryType string,
	source string,
) (*store.Memory, error) {
	if memoryType == "" {
		memoryType = "feedback"
	}

	o := store.NewSaveOption()
	o.Name = name
	o.Content = content
	o.Description = description
	o.Type = memoryType
	o.Source = source
	identifier, e := s.store.CreateMemory(o)

	if e != nil {
		return nil, e
	}

	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	if e = s.indexer.Push(
		fmt.Sprintf("memory/%d", m.Identifier),
		m.Content,
		map[string]string{
			"memory_id": fmt.Sprintf("%d", m.Identifier),
			"type":      m.Type,
		},
	); e != nil {
		return nil, e
	}

	return m, nil
}
