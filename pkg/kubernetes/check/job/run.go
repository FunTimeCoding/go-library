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
	pflag.Bool(RenovateArgument, false, "Run Renovate job")
	pflag.Bool(argument.Wait, false, "Wait for job to complete")
	argument.ParseBind()
	trivy := viper.GetBool(TrivyArgument)
	renovate := viper.GetBool(RenovateArgument)
	wait := viper.GetBool(argument.Wait)
	k := client.NewEnvironment()

	if !trivy && !renovate {
		trivy = true
		renovate = true
	}

	if trivy {
		start(k, wait, constant.TrivyNamespace, constant.TrivyCron)
	}

	if renovate {
		start(k, wait, constant.RenovateNamespace, constant.RenovateCron)
	}
}
