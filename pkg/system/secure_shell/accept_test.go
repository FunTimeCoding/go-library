package secure_shell

import (
	"github.com/funtimecoding/go-library/pkg/system/secure_shell/mock_listener"
	"testing"
)

func TestAccept(t *testing.T) {
	Accept(mock_listener.New())
}
