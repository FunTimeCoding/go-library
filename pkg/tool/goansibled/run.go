package goansibled

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/option"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/runner"
)

func Run(
	o *option.Ansible,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	lifecycle.New(
		l,
		lifecycle.WithWorker(runner.New(o, l, r)),
	).RunUntilSignal()
}
