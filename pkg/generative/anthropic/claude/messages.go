package claude

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/message"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"os"
	"path/filepath"
)

func (c *Client) Messages(sessionIdentifier string) []message.Message {
	path := filepath.Join(
		c.base,
		join.Empty(sessionIdentifier, constant.NotationLogExtension),
	)
	f, e := os.Open(path)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(f)
	var result []message.Message
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	for scanner.Scan() {
		var line notation.Line

		if json.Unmarshal(scanner.Bytes(), &line) != nil {
			continue
		}

		if line.Type != "user" && line.Type != "assistant" {
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

		if text == "" {
			continue
		}

		result = append(
			result,
			message.Message{
				Role:      m.Role,
				Text:      text,
				Timestamp: line.Timestamp,
				IsMeta:    line.Meta || isSystemNoise(text),
			},
		)
	}

	return result
}
