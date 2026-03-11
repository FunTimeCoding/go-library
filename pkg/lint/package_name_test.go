package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	library "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestPackageNameBlacklisted(t *testing.T) {
	l := PackageName(
		library.Alfa,
		strings.NewReader("package api\n"),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "package_name",
				Text:     "Blacklisted package name",
				Path:     "Alfa",
				Line:     1,
				LineText: "package api",
			},
		},
		"",
		l,
	)
}

func TestPackageNameAllowed(t *testing.T) {
	l := PackageName(
		library.Bravo,
		strings.NewReader("package server\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}
