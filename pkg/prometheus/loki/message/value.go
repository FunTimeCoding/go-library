package message

import "fmt"

func (m *Message) Value(key string) string {
	if m.Values == nil {
		return ""
	}

	if v, okay := m.Values[key]; okay {
		return fmt.Sprintf("%s", v)
	}

	return ""
}
