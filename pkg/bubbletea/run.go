package bubbletea

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	library "github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
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
				path = system.Join(
					system.Home(),
					systemConstant.MacOSLibrary,
					systemConstant.MacOsLogs,
					constant.LogFile,
				)
			case systemConstant.Linux:
				path = system.Join("/tmp", constant.LogFile)
			default:
				path = constant.LogFile
			}
		} else {
			system.EnsurePathExists(systemConstant.Temporary)
			path = system.Join(systemConstant.Temporary, constant.LogFile)
		}

		f := Log(path)
		defer errors.LogClose(f)
	}

	RunProgram(tea.NewProgram(m))
}
