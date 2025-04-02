package generate

import "fmt"

func (g *Generate) Print() {
	fmt.Printf("Total: %dms\n", g.Total)
	fmt.Printf("  Load: %dms\n", g.Load)
	fmt.Printf("Prompt tokens: %.0f/s\n", g.PromptTokens)
	fmt.Printf("  Evaluation: %dms\n", g.PromptEvaluation)
	fmt.Printf("Tokens: %.0f/s\n", g.Tokens)
	fmt.Printf("  Evaluation: %dms\n", g.Evaluation)
}
