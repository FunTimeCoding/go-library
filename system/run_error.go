package system

import "os/exec"

func RunError(s ...string) (string, error) {
	result, e := exec.Command(s[0], s[1:]...).CombinedOutput()

	return string(result), e
}
