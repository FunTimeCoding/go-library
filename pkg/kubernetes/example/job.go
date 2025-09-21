package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/kubernetes/filter"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
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
	deletePreviousManualJob(k, constant.RenovateNamespace)
	j := k.CreateJobFromCron(
		constant.RenovateNamespace,
		constant.RenovateCron,
		constant.ManualCron,
	)
	fmt.Printf("Job: %s\n", j.Name)

	if wait {
		waitForFinish(k, constant.RenovateNamespace, j.Name)
		printJobLog(k, constant.RenovateNamespace, j.Name)
	}
}

func trivy(
	k *client.Client,
	wait bool,
) {
	deletePreviousManualJob(k, constant.TrivyNamespace)
	j := k.CreateJobFromCron(
		constant.TrivyNamespace,
		constant.TrivyCron,
		constant.ManualCron,
	)
	fmt.Printf("Job: %s\n", j.Name)

	if wait {
		waitForFinish(k, constant.TrivyNamespace, j.Name)
		printJobLog(k, constant.TrivyNamespace, j.Name)
	}
}

func printJobs(
	k *client.Client,
	namespace string,
) {
	fmt.Printf("Jobs in %s:\n", namespace)

	for _, j := range k.Jobs(namespace) {
		fmt.Printf("  Job: %s\n", j.Name)
	}
}

func deletePreviousManualJob(
	k *client.Client,
	namespace string,
) {
	if j := k.Job(
		namespace,
		constant.ManualCron,
	); j != nil {
		fmt.Printf("Delete job: %s\n", j.Name)
		k.DeleteJobWatch(namespace, j.Name)
	}
}

func waitForFinish(
	k *client.Client,
	namespace string,
	job string,
) {
	time.Sleep(10 * time.Second)
	printJobs(k, namespace)
	errors.PanicOnError(k.WaitForJob(namespace, job, 0))
}

func printJobLog(
	k *client.Client,
	namespace string,
	job string,
) {
	j := k.Job(namespace, job)

	if j == nil {
		fmt.Printf("Not found: %s/%s\n", namespace, job)

		return
	}

	if len(j.Raw.Status.Conditions) > 0 {
		for _, p := range j.Raw.Status.Conditions {
			if p.Type != "Complete" {
				continue
			}

			if p.Status != "True" {
				continue
			}

			t := j.Raw.Status.CompletionTime.String()
			fmt.Printf("Completed: %s\n", t)
		}
	}

	if j.Raw.Status.Active > 0 {
		fmt.Printf("Active: %d\n", j.Raw.Status.Active)
	}

	if j.Raw.Status.Failed > 0 {
		fmt.Printf("Failed: %d\n", j.Raw.Status.Failed)
	}

	if j.Raw.Status.Succeeded > 0 {
		fmt.Printf("Succeeded: %d\n", j.Raw.Status.Succeeded)
	}

	for _, p := range k.Pods(
		filter.New().AddNamespaces(namespace).AddNames(j.Name),
	) {
		fmt.Printf("Pod: %s\n", p.Name)
		log := k.Log(namespace, p.Name, "")
		fmt.Println("Log:")
		fmt.Println(log)
		fmt.Println("-----")
	}
}
