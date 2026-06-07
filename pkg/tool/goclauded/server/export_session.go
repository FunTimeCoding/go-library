package server

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/writer"
	"os"
	"path/filepath"
	"strings"
)

func (s *Server) exportSession(
	i *session.Session,
	basePath string,
) string {
	o, f := s.service.GetSession(i.Identifier)
	errors.PanicOnError(f)
	alias := ""

	if o != nil {
		alias = o.AliasValue()
	}

	messages := s.service.Messages(i.Identifier)
	created := parseTimestamp(i.Timestamp)
	d := filepath.Join(
		basePath,
		created.Format("2006"),
		created.Format("01"),
	)
	slug := alias

	if slug == "" {
		slug = i.Slug
	}

	if slug == "" {
		slug = i.Identifier[:8]
	}

	name := fmt.Sprintf(
		"%s-%s.md",
		created.Format("2006-01-02"),
		strings.ReplaceAll(
			strings.ToLower(slug),
			" ",
			"-",
		),
	)
	var b strings.Builder
	writer.Print(&b, "---\n")
	writer.Print(&b, "session: %s\n", i.Identifier)

	if alias != "" {
		writer.Print(&b, "alias: %s\n", alias)
	}

	if i.Slug != "" {
		writer.Print(&b, "slug: %s\n", i.Slug)
	}

	writer.Print(
		&b,
		"created: %s\n",
		created.Format("2006-01-02"),
	)
	writer.Print(&b, "---\n\n")

	for _, m := range messages {
		if m.Text == "" || m.IsMeta {
			continue
		}

		writer.Print(&b, "## %s\n\n%s\n\n", m.Role, m.Text)
	}

	system.MakeDirectory(d)
	path := filepath.Join(d, name)
	e := os.WriteFile(path, []byte(b.String()), 0644)

	if e != nil {
		panic(e)
	}

	return path
}
