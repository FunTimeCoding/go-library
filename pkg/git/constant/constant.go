package constant

const (
	VersionPrefix = "v"

	Depth int = 2 // Test directory depth from repository root

	OriginRemote = "origin"

	Directory = ".git"

	MainBranch   = "main"
	MasterBranch = "master"
)

const (
	GitHubHost = "github.com"
	GitLabHost = "gitlab.com"
)

var MainBranches = []string{MainBranch, MasterBranch}
