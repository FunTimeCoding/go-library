package generate_response

import "github.com/ollama/ollama/api"

func New(r *api.GenerateResponse) *Response {
	return &Response{
		Text:             r.Response,
		Total:            r.TotalDuration.Milliseconds(),
		Load:             r.LoadDuration.Milliseconds(),
		PromptEvaluation: r.PromptEvalDuration.Milliseconds(),
		Evaluation:       r.EvalDuration.Milliseconds(),
		PromptTokens:     float64(r.PromptEvalCount) / r.PromptEvalDuration.Seconds(),
		Tokens:           float64(r.EvalCount) / r.EvalDuration.Seconds(),
	}
}
