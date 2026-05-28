package run

import "io"

func (r *Run) Writers(
	stdout io.Writer,
	stderr io.Writer,
) {
	r.stdout = stdout
	r.stderr = stderr
}
