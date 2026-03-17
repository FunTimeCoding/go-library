package message

import "github.com/anthropics/anthropic-sdk-go"

func (m *Message) ToParameterCached(p anthropic.CacheControlEphemeralParam) anthropic.MessageParam {
	return anthropic.MessageParam{
		Role: anthropic.MessageParamRole(m.Role),
		Content: []anthropic.ContentBlockParamUnion{
			{
				OfText: &anthropic.TextBlockParam{
					Text:         m.Content,
					CacheControl: p,
				},
			},
		},
	}
}
