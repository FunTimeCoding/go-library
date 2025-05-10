package chat_response

import "github.com/ollama/ollama/api"

func New(r *api.ChatResponse) *Response {
	return &Response{
		Message:          r.Message,
		Total:            r.TotalDuration.Milliseconds(),
		Load:             r.LoadDuration.Milliseconds(),
		PromptEvaluation: r.PromptEvalDuration.Milliseconds(),
		Evaluation:       r.EvalDuration.Milliseconds(),
		PromptTokens:     float64(r.PromptEvalCount) / r.PromptEvalDuration.Seconds(),
		Tokens:           float64(r.EvalCount) / r.EvalDuration.Seconds(),

		Raw: r,
	}
}
