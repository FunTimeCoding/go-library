package claude

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/session"
	"os"
	"path/filepath"
	"strings"
)

func scanSession(path string) *session.Session {
	f, e := os.Open(path)

	if e != nil {
		return session.Stub()
	}

	defer errors.PanicClose(f)
	identifier := strings.TrimSuffix(
		filepath.Base(path),
		constant.NotationLogExtension,
	)
	s := session.New(identifier)
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	filled := false

	for scanner.Scan() {
		s.Lines++

		if filled {
			continue
		}

		var line notation.Line

		if json.Unmarshal(scanner.Bytes(), &line) != nil {
			continue
		}

		if line.Slug != "" && s.Slug == "" {
			s.Slug = line.Slug
		}

		if line.Timestamp != "" && s.Timestamp == "" {
			s.Timestamp = line.Timestamp
		}

		if line.CWD != "" && s.CWD == "" {
			s.CWD = line.CWD
		}

		if line.GitBranch != "" && s.Branch == "" {
			s.Branch = line.GitBranch
		}

		if s.Slug != "" && s.Timestamp != "" {
			filled = true
		}
	}

	return s
}
