package chroma

import (
	"context"
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
)

type Client struct {
	context context.Context
	client  v2.Client
}
