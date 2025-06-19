package protocol

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/chromium"
)

type Protocol struct {
	client  *chromium.Client
	context context.Context
}
