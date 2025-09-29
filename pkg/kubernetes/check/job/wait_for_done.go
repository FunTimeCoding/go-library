package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"time"
)

func waitForDone(
	k *client.Client,
	namespace string,
	job string,
) {
	fmt.Println("Sleep before wait")
	time.Sleep(10 * time.Second)

	printJobs(k, namespace)

	fmt.Println("Wait for job")
	errors.PanicOnError(k.WaitForJob(namespace, job, 0))
}
