package client

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/google/uuid"
)

func generateVerifier() string {
	return join.Empty(uuid.New().String(), uuid.New().String())
}
