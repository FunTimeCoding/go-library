package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/constant"
	"strings"
)

func (s *Service) Profile(topic string) (*ProfileResult, error) {
	always, e := s.ListMemoriesWithContent(constant.AlwaysTag)

	if e != nil {
		return nil, fmt.Errorf("%w: %v", ErrorAlwaysLoad, e)
	}

	result := &ProfileResult{Always: always}

	if topic != "" {
		exclude := make([]string, len(always))

		for i, m := range always {
			exclude[i] = fmt.Sprintf("memory/%d", m.Identifier)
		}

		relevant, f := s.SearchRelevant(topic, 10, exclude)

		if f != nil {
			return nil, fmt.Errorf("%w: %v", ErrorRelevantSearch, f)
		}

		result.Relevant = relevant
	}

	alwaysIDs := map[int64]bool{}

	for _, m := range always {
		alwaysIDs[m.Identifier] = true
	}

	relevantIDs := map[int64]bool{}

	for _, r := range result.Relevant {
		relevantIDs[r.Identifier] = true
	}

	allMemories, e := s.ListMemories("", "", true)

	if e != nil {
		return nil, fmt.Errorf("%w: %v", ErrorMemoryList, e)
	}

	for _, m := range allMemories {
		if !alwaysIDs[m.Identifier] && !relevantIDs[m.Identifier] {
			result.Index = append(result.Index, m)
		}
	}

	impressions, e := s.LatestImpressions(10)

	if e == nil {
		result.Impressions = impressions
	}

	completions, e := s.ListCompletions()

	if e == nil {
		for _, r := range completions {
			name := r.Path

			if i := strings.LastIndex(name, separator.Slash); i >= 0 {
				name = name[:i]
			}

			result.Completions = append(
				result.Completions,
				CompletionEntry{SessionName: name, Body: r.Body},
			)
		}
	}

	return result, nil
}
