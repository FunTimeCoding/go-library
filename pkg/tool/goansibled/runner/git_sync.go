package runner

func (r *Runner) gitSync() bool {
	r.gitFetch()
	local := r.gitRevision("HEAD")
	remote := r.gitRevision(RemoteBranch)

	if local == remote {
		r.logger.Structured("git_sync", "status", "unchanged")

		return false
	}

	r.logger.Structured(
		"git_sync",
		"status",
		"changed",
		"local",
		local,
		"remote",
		remote,
	)
	r.gitPull()

	return true
}
