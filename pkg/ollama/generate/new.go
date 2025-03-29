package generate

import "github.com/ollama/ollama/api"

func New(r *api.GenerateResponse) *Generate {
	return &Generate{
		Text:             r.Response,
		Total:            r.TotalDuration.Milliseconds(),
		Load:             r.LoadDuration.Milliseconds(),
		PromptEvaluation: r.PromptEvalDuration.Milliseconds(),
		Evaluation:       r.EvalDuration.Milliseconds(),
		Tokens:           float64(r.EvalCount) / r.EvalDuration.Seconds(),
	}
}
