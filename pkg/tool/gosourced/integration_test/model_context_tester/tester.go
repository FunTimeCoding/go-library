package model_context_tester

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context_client"
	"github.com/funtimecoding/go-library/pkg/tool/gosourced/integration_test/base"
)

type Tester struct {
	server    *base.Server
	Client    *model_context_client.Client
	Directory string
}
