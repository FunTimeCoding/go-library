package prompts

import "github.com/funtimecoding/go-library/pkg/ollama/prompt"

func ClassifyAlert() *prompt.Prompt {
	p := prompt.New()
	p.Instructions(
		"Decide if this Prometheus alert is already-broken or not-yet-broken",
	)
	p.ToClassify("ProbeFail")
	p.AnswerFormat("JSON, 2 string keys: Reason, Answer")
	p.AnswerChoice("Reason", "One concise reasoning sentence.")
	p.AnswerChoice(
		"Answer",
		"Choice between already-broken, not-yet-broken",
	)
	p.Example("already-broken", []string{"DiskFull"})
	p.Example(
		"not-yet-broken",
		[]string{"DiskNearFull", "LatencyHigh"},
	)

	return p
}
