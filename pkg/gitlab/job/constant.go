package job

const (
	NoUser    = "no user"
	NoProject = "no project"

	FailConcern = "fail"
	Timeout     = "timeout"

	traceTimeoutMatch = "ERROR: Job failed (system failure): prepare environment: waiting for pod running: timed out waiting for pod to start."
)
