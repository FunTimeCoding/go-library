package runner

func (r *Runner) CurrentHead() string {
	var head string
	r.recovery.Run(func() { head = r.gitRevision("HEAD") })

	return head
}
