package classify_prompt

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (p *Prompt) Example(
	choice string,
	v []string,
) {
	p.examples = append(
		p.examples,
		fmt.Sprintf("%s: %s", choice, join.Comma(v)),
	)
}
