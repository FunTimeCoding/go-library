package service

import "fmt"

func (s *Service) ForgetMemory(
	identifier int64,
	source string,
) error {
	e := s.store.ForgetMemory(identifier, source)

	if e != nil {
		return e
	}

	return s.indexer.Delete(fmt.Sprintf("memory/%d", identifier))
}
