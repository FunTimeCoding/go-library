package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	TrivyArgument    = "trivy"
	RenovateArgument = "renovate"
)

func Job() {
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

func renovate(
	k *client.Client,
	wait bool,
) {
	for _, j := range k.CronJobs(constant.RenovateNamespace) {
		fmt.Printf("Delete: %s\n", j.Name)
		k.DeleteJobWatch(constant.RenovateNamespace, j.Name)
	}

	if t := k.CronJob(constant.RenovateNamespace, "missing"); t != nil {
		fmt.Printf("Job: %s\n", t.Name)
	}

	k.DeleteJobWatch(constant.RenovateNamespace, constant.ManualCron)
	j := k.CreateJobFromCron(
		constant.RenovateNamespace,
		constant.RenovateCron,
		constant.ManualCron,
	)
	fmt.Printf("Job: %s\n", j.Name)

	if wait {
		errors.PanicOnError(
			k.WaitForJob(
				constant.RenovateNamespace,
				constant.ManualCron,
				0,
			),
		)
	}
}

func trivy(
	k *client.Client,
	wait bool,
) {
	for _, j := range k.CronJobs(constant.TrivyNamespace) {
		fmt.Printf("Delete: %s\n", j.Name)
		k.DeleteJobWatch(constant.TrivyNamespace, j.Name)
	}

	if t := k.CronJob(constant.TrivyNamespace, "missing"); t != nil {
		fmt.Printf("Job: %s\n", t.Name)
	}

	k.DeleteJobWatch(constant.TrivyNamespace, constant.ManualCron)
	j := k.CreateJobFromCron(
		constant.TrivyNamespace,
		constant.TrivyCron,
		constant.ManualCron,
	)
	fmt.Printf("Job: %s\n", j.Name)

	if wait {
		errors.PanicOnError(
			k.WaitForJob(
				constant.TrivyNamespace,
				constant.ManualCron,
				0,
			),
		)
	}
}
