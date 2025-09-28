package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

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
