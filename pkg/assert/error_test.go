package assert

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	Error(t, fmt.Errorf("something went wrong"))
}
