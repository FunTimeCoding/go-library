package job

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Run() {
	pflag.Bool(TrivyArgument, false, "Run Trivy job")
	pflag.Bool(LabArgument, false, "Renovate GitLab")
	pflag.Bool(HubArgument, false, "Renovate GitHub")
	pflag.Bool(argument.Wait, false, "Wait for job to complete")
	argument.ParseBind()
	trivy := viper.GetBool(TrivyArgument)
	lab := viper.GetBool(LabArgument)
	hub := viper.GetBool(HubArgument)
	wait := viper.GetBool(argument.Wait)
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
