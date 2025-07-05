package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	VersionPrefix = "v"

	Depth int = 2 // Test directory depth from repository root

	OriginRemote = "origin"

	Directory = ".git"

	MainBranch   = "main"
	MasterBranch = "master"

	GitHubHost = "github.com"
	GitLabHost = "gitlab.com"

	Command = "git"
	Tag     = "tag"
	Clone   = "clone"
	Status  = "status"

	RevParse = "rev-parse"
	GitDir   = "--git-dir"

	Porcelain = "--porcelain"

	Fetch     = "fetch"
	Prune     = "--prune"
	PruneTags = "--prune-tags"

	Push = "push"
	Tags = "--tags"
)

var (
	MainBranches = []string{MainBranch, MasterBranch}

	Format = option.ExtendedColor.Copy()
)
