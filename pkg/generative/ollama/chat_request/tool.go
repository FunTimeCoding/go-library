package chat_request

import "github.com/ollama/ollama/api"

func (r *Request) Tool(
	toolType string,
	items any,
	f api.ToolFunction,
) {
	r.request.Tools = append(
		r.request.Tools,
		api.Tool{Type: toolType, Items: items, Function: f},
	)
}
