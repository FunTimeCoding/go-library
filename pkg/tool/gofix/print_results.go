package gofix

import "github.com/funtimecoding/go-library/pkg/lint/output"

func printResults(entries []result, summary bool) bool {
	return output.PrintResults(entries, summary)
}
