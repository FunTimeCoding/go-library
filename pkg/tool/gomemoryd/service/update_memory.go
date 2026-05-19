package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (s *Service) UpdateMemory(
	identifier int64,
	name string,
	content string,
	description string,
	source string,
) (*store.Memory, error) {
	existing, e := s.Store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	o := store.NewSaveOption()
	o.Name = name
	o.Content = content
	o.Description = description
	o.Tags = existing.Tags
	o.Source = source
	e = s.Store.UpdateMemory(identifier, o)

	if e != nil {
		return nil, e
	}

	m, e := s.Store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	if e = s.indexer.Push(
		fmt.Sprintf("memory/%d", m.Identifier),
		m.Content,
	); e != nil {
		return nil, e
	}

	return m, nil
}
