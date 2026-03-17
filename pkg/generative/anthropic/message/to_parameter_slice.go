package message

import "github.com/anthropics/anthropic-sdk-go"

func ToParameterSlice(v []*Message) []anthropic.MessageParam {
	result := make([]anthropic.MessageParam, len(v))

	for i, m := range v {
		result[i] = m.ToParameter()
	}

	return result
}
