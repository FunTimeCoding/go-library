package message

import "github.com/anthropics/anthropic-sdk-go"

func ToParamSlice(messages []*Message) []anthropic.MessageParam {
	result := make([]anthropic.MessageParam, len(messages))

	for i, m := range messages {
		result[i] = m.ToParam()
	}

	return result
}
