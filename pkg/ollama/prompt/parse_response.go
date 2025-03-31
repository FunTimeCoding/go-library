package prompt

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (p *Prompt) ParseResponse(s string) string {
	var result *Response
	notation.DecodeStrict(s, &result, true)
	result.Classified = p.toClassify

	return fmt.Sprintf("%s\n", notation.Encode(result, true))
}
