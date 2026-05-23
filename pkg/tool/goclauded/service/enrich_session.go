package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) EnrichSession(identifier string) {
	e := s.client.Resolve(identifier)

	if e.Identifier == "" {
		return
	}

	updates := map[string]any{}

	if e.Slug != "" {
		updates[constant.Slug] = e.Slug
	}

	if e.WorkDirectory != "" {
		updates["work_directory"] = e.WorkDirectory
	}

	if e.Branch != "" {
		updates["branch"] = e.Branch
	}

	if e.Timestamp != "" {
		updates["session_timestamp"] = e.Timestamp
	}

	if e.Lines > 0 {
		updates["lines"] = e.Lines
	}

	peek := s.client.Peek(identifier)

	if len(peek.UserMessages) > 0 {
		updates["turn_count"] = len(peek.UserMessages)
		firstMessage := peek.UserMessages[0]

		if len(firstMessage) > 80 {
			firstMessage = firstMessage[:80]
		}

		updates["first_message"] = firstMessage
	}

	if len(updates) > 0 {
		s.store.UpdateFields(identifier, updates)
	}
}
