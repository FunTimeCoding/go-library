package constant

import (
	"github.com/funtimecoding/go-library/pkg/monitor/collector"
	"github.com/funtimecoding/go-library/pkg/monitor/helper"
	"github.com/funtimecoding/go-library/pkg/monitor/item/source"
	"time"
)

var (
	DefaultWeights = helper.SeverityWeights(3, 2, 1)

	DefaultSource = source.New(
		1,
		0,
		0,
		0,
		0,
		DefaultWeights,
	)

	MonitorCollector = collector.New(
		"undefined",
		"monitor",
		"undefined",
		0,
		DefaultSource,
	) // For monitor errors

	GoAlert = collector.New(
		"goalert",
		"alert",
		"Prometheus alerts",
		2*time.Minute,
		DefaultSource,
	)
	GoContainer = collector.New(
		"gocontainer",
		"podman-container",
		"Podman containers",
		5*time.Minute,
		DefaultSource,
	)
	GoFile = collector.New(
		"gofile",
		"file",
		"File changes",
		time.Minute,
		DefaultSource,
	)
	GoGenie = collector.New(
		"gogenie",
		"opsgenie",
		"Opsgenie alerts",
		time.Minute,
		DefaultSource,
	)
	GoGitHub = collector.New(
		"gogithub",
		"github",
		"GitHub jobs",
		5*time.Minute,
		DefaultSource,
	)
	GoGitLab = collector.New(
		"gogitlab",
		"gitlab",
		"GitLab jobs",
		5*time.Minute,
		DefaultSource,
	)
	GoGitStatus = collector.New(
		"gogitstatus",
		"git",
		"Git repositories",
		10*time.Minute,
		DefaultSource,
	)
	GoImage = collector.New(
		"goimage",
		"podman-image",
		"Podman images",
		10*time.Minute,
		DefaultSource,
	)
	GoJira = collector.New(
		"gojira",
		"jira",
		"Jira issues",
		5*time.Minute,
		DefaultSource,
	)
	GoKevt = collector.New(
		"gokevt",
		"event",
		"Kubernetes events",
		5*time.Minute,
		DefaultSource,
	)
	GoSentry = collector.New(
		"gosentry",
		"sentry",
		"Sentry issues",
		time.Minute,
		DefaultSource,
	)
	GoSilence = collector.New(
		"gosilence",
		"silence",
		"silences",
		time.Minute,
		DefaultSource,
	)
	GoVersion = collector.New(
		"goversion",
		"go",
		"Go projects",
		10*time.Minute,
		DefaultSource,
	)
	GoVolume = collector.New(
		"govolume",
		"podman-volume",
		"Podman volumes",
		10*time.Minute,
		DefaultSource,
	)

	Collectors = []*collector.Collector{
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
