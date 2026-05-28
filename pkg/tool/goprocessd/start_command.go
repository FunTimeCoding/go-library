package goprocessd

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/procfile"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/server"
	"github.com/spf13/cobra"
	"os"
)

func startCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start [PROCESS...]",
		Short: "Start all or specified processes",
		RunE: func(
			cmd *cobra.Command,
			arguments []string,
		) error {
			procfilePath, e := cmd.Flags().GetString("file")
			errors.PanicOnError(e)
			envrcPath, e := cmd.Flags().GetString("envrc")
			errors.PanicOnError(e)
			entries, e := procfile.Parse(procfilePath)

			if e != nil {
				return e
			}

			if len(arguments) > 0 {
				entries = filterEntries(entries, arguments)

				if len(entries) == 0 {
					return fmt.Errorf("no matching processes found")
				}
			}

			env := environment.New(os.Environ())

			if e := env.Load(envrcPath); e != nil {
				errors.Printf("warning: %s\n", e)
			}

			s := server.New(
				entries,
				env,
				procfilePath,
				envrcPath,
				socketPath(procfilePath),
			)

			return s.Run()
		},
	}
}
