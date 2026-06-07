package service

import "github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"

func (s *Service) Summary(
	since string,
	until string,
	groupBy string,
) ([]store.SummaryRow, error) {
	return s.store.Summary(since, until, groupBy)
}
