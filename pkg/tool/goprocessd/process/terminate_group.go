package process

import (
	"golang.org/x/sys/unix"
	"os"
)

func terminateGroup(pid int) error {
	pgid, e := unix.Getpgid(pid)

	if e != nil {
		return e
	}

	target := pid

	if pgid == pid {
		target = -1 * pid
	}

	p, e := os.FindProcess(target)

	if e != nil {
		return e
	}

	return p.Signal(os.Interrupt)
}
