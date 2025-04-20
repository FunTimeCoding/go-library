package generate_response

import "fmt"

func (r *Response) Format() {
	fmt.Printf("Total: %dms\n", r.Total)
	fmt.Printf("  Load: %dms\n", r.Load)
	fmt.Printf("  Prompt evaluation: %dms\n", r.PromptEvaluation)
	fmt.Printf("    Tokens/s: %.0f\n", r.PromptTokens)
	fmt.Printf("  Evaluation: %dms\n", r.Evaluation)
	fmt.Printf("    Tokens/s: %.0f\n", r.Tokens)
}
