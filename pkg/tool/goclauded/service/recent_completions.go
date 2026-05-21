package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"

func (s *Service) RecentCompletions() []completion.Completion {
	return s.store.RecentCompletions()
}
