package client

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa, upper.Bravo, net.IP{}, nil))
}
