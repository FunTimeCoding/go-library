package code

import "github.com/google/go-github/v78/github"

type Code struct {
	Hash string
	Name string
	Path string

	Raw *github.CodeResult
}
