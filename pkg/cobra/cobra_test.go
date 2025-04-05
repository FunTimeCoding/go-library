package cobra

import "testing"

func TestCobra(t *testing.T) {
	Execute(New("ls"))
}
