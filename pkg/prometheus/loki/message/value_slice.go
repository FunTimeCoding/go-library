package message

import "fmt"

func (m *Message) ValueSlice(key string) []string {
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
