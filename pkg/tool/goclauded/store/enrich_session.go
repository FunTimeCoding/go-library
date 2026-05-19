package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) EnrichSession(
	identifier string,
	c *claude.Client,
) {
	e := c.Resolve(identifier)

	if e.Identifier == "" {
		return
	}

	updates := map[string]any{}

	if e.Slug != "" {
		updates["slug"] = e.Slug
	}

	firstMessage := c.FirstUserMessage(identifier)

	if firstMessage != "" {
		if len(firstMessage) > 80 {
			firstMessage = firstMessage[:80]
		}

		updates["first_message"] = firstMessage
	}

	peek := c.Peek(identifier)

	if len(peek.UserMessages) > 0 {
		updates["turn_count"] = len(peek.UserMessages)
	}

	if len(updates) > 0 {
		errors.PanicOnError(
			s.database.Model(session.New()).
				Where("identifier = ?", identifier).
				Updates(updates).Error,
		)
	}
}
