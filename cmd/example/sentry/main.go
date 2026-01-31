package main

import "github.com/funtimecoding/go-library/pkg/errors/sentry/example"

func main() {
	example.Issue()

	if false {
		example.Capture()
		example.TrackedIssue()
		example.Organization()
		example.Project()
		example.Star()
	}
}
