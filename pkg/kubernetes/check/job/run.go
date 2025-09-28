package job

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Run() {
	pflag.Bool(TrivyArgument, false, "Run Trivy job")
	pflag.Bool(RenovateArgument, false, "Run Renovate job")
	pflag.Bool(argument.Wait, false, "Wait for job to complete")
	argument.ParseBind()
	runTrivy := viper.GetBool(TrivyArgument)
	runRenovate := viper.GetBool(RenovateArgument)
	wait := viper.GetBool(argument.Wait)
	k := client.NewEnvironment()

	if !runTrivy && !runRenovate {
		runTrivy = true
		runRenovate = true
	}

	if runTrivy {
		trivy(k, wait)

	}

	if runRenovate {
		renovate(k, wait)
	}
}
