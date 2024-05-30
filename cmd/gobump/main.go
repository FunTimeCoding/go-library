package main

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

const major = "major"
const minor = "minor"
const patch = "patch"

var parts = []string{
	major,
	minor,
	patch,
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(
			"Usage: %s INCREASE\nIncrease options: %s\n",
			join.CommaSpace(parts),
			os.Args[0],
		)

		os.Exit(1)
	}

	d := system.WorkingDirectory()

	if !git.IsCleanCommand() {
		s := git.Status(d)

		if !git.IsClean(s, false) {
			fmt.Printf("Not clean:\n%s\n", s.String())

			os.Exit(1)
		}
	}

	git.Fetch()
	versions := git.Versions(d)
	var next *semver.Version

	if len(versions) == 0 {
		next = semver.New("0.0.0")
	} else {
		next = semver.New(versions[len(versions)-1].String())
	}

	increase := os.Args[1]

	switch increase {
	case patch:
		next.BumpPatch()
	case minor:
		next.BumpMinor()
	case major:
		next.BumpMajor()
	default:
		fmt.Printf("Unexpected increase: %s", increase)

		os.Exit(1)
	}

	nextString := fmt.Sprintf(
		"%s%s",
		constant.VersionPrefix,
		next.String(),
	)
	fmt.Printf("Tag: %s\n", nextString)
	git.Tag(nextString)
	git.Push()
}
