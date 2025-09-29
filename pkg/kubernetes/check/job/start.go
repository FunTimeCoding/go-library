package job

import "github.com/funtimecoding/go-library/pkg/kubernetes/client"

func start(
	k *client.Client,
	wait bool,
	namespace string,
	cron string,
) {
	deleteManual(k, namespace)
	j := createManual(k, namespace, cron)

	if wait {
		waitForDone(k, namespace, j.Name)
		printLog(k, namespace, j.Name)
	}
}
