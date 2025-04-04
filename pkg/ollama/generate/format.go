package generate

import "fmt"

func (g *Generate) Print() {
	fmt.Printf("Total: %dms\n", g.Total)
	fmt.Printf("  Load: %dms\n", g.Load)
	fmt.Printf("  Prompt evaluation: %dms\n", g.PromptEvaluation)
	fmt.Printf("    Tokens/s: %.0f\n", g.PromptTokens)
	fmt.Printf("  Evaluation: %dms\n", g.Evaluation)
	fmt.Printf("    Tokens/s: %.0f\n", g.Tokens)
}
