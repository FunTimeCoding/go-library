package message

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (m *Message) FormatChat() string {
	return fmt.Sprintf(
		"%s %s: %s",
		m.Create.Format(time.DateMinute),
		m.From,
		m.Text,
	)
}
