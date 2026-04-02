package main

import "github.com/funtimecoding/go-library/pkg/errors/sentry/example"

func main() {
	example.Whoami()

	if false {
		example.Capture()
		example.Issue()
		example.TrackedIssue()
		example.Organization()
		example.Project()
		example.Star()
	}
}
