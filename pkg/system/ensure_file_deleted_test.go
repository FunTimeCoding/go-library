package system

import "testing"

func TestEnsureFileDeleted(t *testing.T) {
	EnsureFileDeleted("neverExistingFile")
}
