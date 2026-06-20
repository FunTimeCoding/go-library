package runner

func (r *Runner) syncWithDiff() *SyncResult {
	r.gitFetch()
	local := r.gitRevision("HEAD")
	remote := r.gitRevision(RemoteBranch)

	if local == remote {
		r.logger.Structured("sync", Status, "unchanged")

		return &SyncResult{}
	}

	r.logger.Structured(
		"sync",
		Status,
		"changed",
		"local",
		local,
		"remote",
		remote,
	)
	diff := r.gitDiffLog(local, remote)
	r.gitPull()

	return &SyncResult{Changed: true, Diff: diff}
}
