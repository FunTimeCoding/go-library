package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

func deleteManual(
	k *client.Client,
	namespace string,
) {
	name := constant.ManualCron
	fmt.Printf("Delete job %s in %s\n", name, namespace)

	if j := k.Job(namespace, name); j != nil {
		fmt.Printf("  %s\n", j.Name)
		k.DeleteJobWatch(namespace, j.Name)
	}
}
