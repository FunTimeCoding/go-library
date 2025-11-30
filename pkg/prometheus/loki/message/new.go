package message

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"
	"github.com/funtimecoding/go-library/pkg/time"
	timeLibrary "time"
)

const (
	TextType     = "text"
	NotationType = "json"
)

type Message struct {
	Time      timeLibrary.Time
	Stream    string
	Type      string
	Text      string
	Values    map[string]any
	RawStream *response.Stream
}

func New(
	e []string,
	r *response.Stream,
) *Message {
	if false {
		fmt.Printf("Value: %s\n", r.Stream)
		fmt.Printf("  Timestamp: %s\n", e[0])
		fmt.Printf("  Line: %s\n", e[1])
	}

	var messageType string
	result := &Message{
		Time:      time.FromUnixNanoString(e[0]),
		Stream:    r.Stream,
		RawStream: r,
	}

	if v, okay := parseNotation(e[1]); okay {
		messageType = NotationType
		result.Values = v
	} else {
		messageType = TextType
		result.Text = e[1]
	}

	result.Type = messageType

	return result
}

func NewSlice(v response.Data) ([]*Message, *Meta) {
	var result []*Message

	for _, e := range v.Result {
		for _, a := range e.Values {
			result = append(result, New(a, &e.Stream))
		}
	}

	return result, &Meta{Type: v.ResultType, Statistic: v.Stats}
}

func (m *Message) ValueString(key string) string {
	if m.Values == nil {
		return ""
	}

	if v, okay := m.Values[key]; okay {
		return fmt.Sprintf("%s", v)
	}

	return ""
}

func (m *Message) ValueInt(key string) int {
	if m.Values == nil {
		return 0
	}

	if v, okay := m.Values[key]; okay {
		return int(v.(float64))
	}

	return 0
}

func (m *Message) ValueStringSlice(key string) []string {
	if m.Values == nil {
		return nil
	}

	if v, okay := m.Values[key]; okay {
		var result []string

		for _, a := range v.([]any) {
			result = append(result, fmt.Sprintf("%s", a))
		}

		return result
	}

	return nil
}

func (m *Message) Value(key string) any {
	if m.Values == nil {
		return nil
	}

	return m.Values[key]
}

type Meta struct {
	Type      string
	Statistic response.Statistic
}

func parseNotation(s string) (map[string]any, bool) {
	var result map[string]any

	if e := json.Unmarshal([]byte(s), &result); e != nil {
		return nil, false
	}

	return result, true
}
