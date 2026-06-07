package assert

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorIs(t *testing.T) {
	sentinel := errors.New("not found")
	ErrorIs(t, fmt.Errorf("thing not found: %w", sentinel), sentinel)
}
