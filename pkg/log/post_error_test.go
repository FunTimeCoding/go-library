package log

import "testing"

func TestPostError(t *testing.T) {
	PostError("https://localhost", "Something went wrong")
}
