package main

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	gitHelper "github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(
			"Usage: %s INCREASE\nIncrease options: major, minor, patch\n",
			os.Args[0],
		)

		os.Exit(1)
	}

	d := system.WorkingDirectory()
	s := gitHelper.Status(d)

	if !gitHelper.IsClean(s, false) {
		fmt.Printf("Status: %s\n", s.String())

		os.Exit(1)
	}

	system.Run("git", "fetch", "--prune", "--prune-tags")
	var versions semver.Versions

	for _, element := range gitHelper.Tags(d) {
		versions = append(versions, semver.New(element))
	}

	semver.Sort(versions)
	var next *semver.Version

	if len(versions) == 0 {
		next = semver.New("0.0.0")
	} else {
		next = semver.New(versions[len(versions)-1].String())
	}

	increase := os.Args[1]

	switch increase {
	case "patch":
		next.BumpPatch()
	case "minor":
		next.BumpMinor()
	case "major":
		next.BumpMajor()
	default:
		fmt.Printf("Unexpected increase: %s", increase)

		os.Exit(1)
	}

	fmt.Printf("Tag: %s\n", next)
	system.Run("git", "tag", next.String())
	system.Run("git", "push", "origin", "--tags")
}
