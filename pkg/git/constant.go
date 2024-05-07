package git

const (
	VersionPrefix = "v"

	depth = int(2) // Test directory depth from repository root

	OriginRemote = "origin"

	Directory = ".git"

	MainBranch   = "main"
	MasterBranch = "master"
)

var MainBranches = []string{MainBranch, MasterBranch}
