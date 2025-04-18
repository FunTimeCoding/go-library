package classify_prompt

type Prompt struct {
	instructions string
	toClassify   string
	answerFormat string
	choices      []string
	examples     []string
}
