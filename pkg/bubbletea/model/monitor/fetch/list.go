package fetch

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"strings"
	"time"
)

type Setting struct {
	Command  string
	Interval time.Duration
}

var (
	GoAlert = &Setting{
		Command:  constant.GoAlert,
		Interval: 2 * time.Minute,
	}
	GoContainer = &Setting{
		Command:  constant.GoContainer,
		Interval: 5 * time.Minute,
	}
	GoVolume = &Setting{
		Command:  constant.GoVolume,
		Interval: 10 * time.Minute,
	}
	GoImage = &Setting{
		Command:  constant.GoImage,
		Interval: 10 * time.Minute,
	}
	GoFile = &Setting{
		Command:  constant.GoFile,
		Interval: time.Minute,
	}
	GoGenie = &Setting{
		Command:  constant.GoGenie,
		Interval: 1 * time.Minute,
	}
	GoGitHub = &Setting{
		Command:  constant.GoGitHub,
		Interval: 5 * time.Minute,
	}
	GoGitLab = &Setting{
		Command:  constant.GoGitLab,
		Interval: 5 * time.Minute,
	}
	GoGitStatus = &Setting{
		Command:  constant.GoGitStatus,
		Interval: 10 * time.Minute,
	}
	GoJira = &Setting{
		Command:  constant.GoJira,
		Interval: 5 * time.Minute,
	}
	GoKevt = &Setting{
		Command:  constant.GoKevt,
		Interval: 5 * time.Minute,
	}
	GoSentry = &Setting{
		Command:  constant.GoSentry,
		Interval: 1 * time.Minute,
	}
	GoSilence = &Setting{
		Command:  constant.GoSilence,
		Interval: 1 * time.Minute,
	}
	GoVersion = &Setting{
		Command:  constant.GoVersion,
		Interval: 10 * time.Minute,
	}
	Settings = []*Setting{
		GoAlert,
		GoContainer,
		GoFile,
		GoGenie,
		GoGitHub,
		GoGitLab,
		GoGitStatus,
		GoImage,
		GoJira,
		GoKevt,
		GoSentry,
		GoSilence,
		GoVersion,
		GoVolume,
	}
)

func List() []string {
	var result []string

	for _, cmd := range Settings {
		result = append(result, cmd.Command)
	}

	if s := environment.GetDefault(
		constant.PluginEnvironment,
		"",
	); s != "" {
		if strings.HasPrefix(s, separator.Plus) {
			result = append(result, split.Comma(s)...)
		} else {
			result = split.Comma(s)
		}
	}

	return result
}
