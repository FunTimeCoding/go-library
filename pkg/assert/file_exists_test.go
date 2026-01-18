package assert

import "testing"

func TestFileExists(t *testing.T) {
	FileExists(t, "./file_exists.go")
}
