package tunnel

import "os/exec"

type Tunnel struct {
	command   *exec.Cmd
	started   chan struct{}
	listening chan struct{}
	stopped   chan struct{}
	Verbose   bool
	NoOutput  bool
	CleanStop bool
}
