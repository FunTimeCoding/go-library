package classify_prompt

import "fmt"

func (p *Prompt) AnswerChoice(
	k string,
	v string,
) {
	p.choices = append(p.choices, fmt.Sprintf("%s: %s", k, v))
}
