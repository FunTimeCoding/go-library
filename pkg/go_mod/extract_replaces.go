package go_mod

import "github.com/funtimecoding/go-library/pkg/expression"

func ExtractReplaces(mod string) string {
	return expression.SubMatchIndex(`(?s)replace (.*)`, mod, 0)
}
