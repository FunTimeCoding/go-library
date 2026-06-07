package goprocessd

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/procfile"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/socket"
	"os"
)

func Run(o *option.Option) {
	entries, e := procfile.Parse(o.ProcfilePath)
	errors.PanicOnError(e)
	env := environment.New(os.Environ())

	if e := env.Load(o.EnvrcPath); e != nil {
		errors.Printf("warning: %s\n", e)
	}

	s := server.New(
		entries,
		env,
		o.ProcfilePath,
		o.EnvrcPath,
		socket.Path(o.ProcfilePath),
	)
	errors.PanicOnError(s.Run())
}
