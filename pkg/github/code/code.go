package code

import "github.com/google/go-github/v70/github"

type Code struct {
	Hash string
	Name string
	Path string

	Raw *github.CodeResult
}
