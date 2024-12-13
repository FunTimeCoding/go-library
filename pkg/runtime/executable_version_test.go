package runtime

import "testing"

func TestExecutableVersion(t *testing.T) {
	v := ExecutableVersion()

	if v.Major < 1 {
		t.Errorf(
			"Expected major version to be at least 1, got %d",
			v.Major,
		)
	}

	if v.Minor < 23 {
		t.Errorf(
			"Expected minor version to be at least 23, got %d",
			v.Minor,
		)
	}
}
