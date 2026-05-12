package job

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

func Run() {
	a := argument.NewSimple("kubernetes-job")
	a.Boolean(TrivyArgument, false, "Run Trivy job")
	a.Boolean(LabArgument, false, "Renovate GitLab")
	a.Boolean(HubArgument, false, "Renovate GitHub")
	a.Boolean(argument.Wait, false, "Wait for job to complete")
	a.ParseSimple()
	trivy := a.GetBoolean(TrivyArgument)
	lab := a.GetBoolean(LabArgument)
	hub := a.GetBoolean(HubArgument)
	wait := a.GetBoolean(argument.Wait)
	k := client.NewEnvironment()

	if !trivy && !lab && !hub {
		trivy = true
		lab = true
		hub = true
	}

	if trivy {
		start(
			k,
			wait,
			constant.TrivyNamespace,
			constant.TrivyCron,
			constant.ManualJob,
		)
	}

	if lab {
		start(
			k,
			wait,
			constant.RenovateNamespace,
			constant.LabCron,
			constant.ManualLabJob,
		)
	}

	if hub {
		start(
			k,
			wait,
			constant.RenovateNamespace,
			constant.HubCron,
			constant.ManualHubJob,
		)
	}
}
