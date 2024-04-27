package errors

import "testing"

func TestPost(t *testing.T) {
	Post("https://localhost", "something went wrong")
}
