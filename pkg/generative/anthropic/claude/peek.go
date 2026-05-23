package claude

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/peek"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"os"
	"path/filepath"
)

func (c *Client) Peek(sessionIdentifier string) *peek.Peek {
	path := filepath.Join(
		c.base,
		join.Empty(sessionIdentifier, constant.NotationLogExtension),
	)
	f, e := os.Open(path)

	if e != nil {
		return peek.Stub()
	}

	defer errors.PanicClose(f)
	result := peek.Stub()
	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 1024*1024), 1024*1024)

	for s.Scan() {
		result.LineCount++
		var line notation.Line

		if json.Unmarshal(s.Bytes(), &line) != nil {
			continue
		}

		if line.Type != "user" || line.Meta {
			continue
		}

		if line.Message == nil {
			continue
		}

		var m notation.Message

		if json.Unmarshal(line.Message, &m) != nil {
			continue
		}

		text := extractText(m.Content)

		if isSystemNoise(text) {
			continue
		}

		clean := cleanContent(text)

		if len(clean) > 20 {
			if len(clean) > 100 {
				clean = clean[:100]
			}

			result.UserMessages = append(result.UserMessages, clean)
		}
	}

	return result
}
