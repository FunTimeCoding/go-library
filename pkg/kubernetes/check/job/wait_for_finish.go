package job

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"time"
)

func waitForFinish(
	k *client.Client,
	namespace string,
	job string,
) {
	time.Sleep(10 * time.Second)
	printJobs(k, namespace)
	errors.PanicOnError(k.WaitForJob(namespace, job, 0))
}
