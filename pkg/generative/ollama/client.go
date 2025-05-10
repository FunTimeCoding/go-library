package ollama

import (
	"context"
	"github.com/ollama/ollama/api"
)

type Client struct {
	context context.Context
	client  *api.Client
}
