package chat_response

import "github.com/ollama/ollama/api"

type Response struct {
	Message          api.Message
	Total            int64
	Load             int64
	PromptEvaluation int64
	Evaluation       int64
	PromptTokens     float64
	Tokens           float64

	Raw *api.ChatResponse
}
