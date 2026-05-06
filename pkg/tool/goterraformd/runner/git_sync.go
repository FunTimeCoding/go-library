package runner

import "github.com/funtimecoding/go-library/pkg/tool/goterraformd/constant"

func (r *Runner) gitSync() bool {
	r.gitFetch()
	local := r.gitRevision("HEAD")
	remote := r.gitRevision(RemoteBranch)

	if local == remote {
		r.logger.Structured("git_sync", constant.Status, "unchanged")

		return false
	}

	r.logger.Structured(
		"git_sync",
		constant.Status,
		"changed",
		"local",
		local,
		"remote",
		remote,
	)
	r.gitPull()

	return true
}
