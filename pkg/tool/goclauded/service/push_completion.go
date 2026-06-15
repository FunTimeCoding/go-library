package service

import "fmt"

func (s *Service) pushCompletion(
	slug string,
	sequence int,
	body string,
	metadata map[string]string,
) error {
	return s.completionIndexer.Push(
		fmt.Sprintf("%s/%d", slug, sequence),
		body,
		metadata,
	)
}
