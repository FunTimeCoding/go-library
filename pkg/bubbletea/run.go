package bubbletea

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	library "github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"runtime"
)

func Run(
	m tea.Model,
	log bool,
) {
	if log {
		var path string

		if library.RunningFromBinary() {
			switch runtime.GOOS {
			case systemConstant.Darwin:
				path = join.Absolute(
					system.Home(),
					systemConstant.MacOSLibrary,
					systemConstant.MacOsLogs,
					constant.LogFile,
				)
			case systemConstant.Linux:
				path = join.Absolute(
					systemConstant.Temporary,
					constant.LogFile,
				)
			default:
				path = constant.LogFile
			}
		} else {
			system.EnsurePathExists(systemConstant.Temporary)
			path = join.Relative(systemConstant.Temporary, constant.LogFile)
		}

		f := Log(path)
		defer errors.LogClose(f)
	}

	RunProgram(tea.NewProgram(m))
}
