package client

import (
	"fmt"
	"time"
)

func (c *Client) WaitForJob(
	namespace string,
	name string,
	timeout time.Duration,
) error {
	start := time.Now()

	for {
		if timeout > 0 && time.Since(start) > timeout {
			return fmt.Errorf("job timeout: %s", name)
		}

		j := c.Job(namespace, name)

		if j == nil {
			return nil
		}

		if j.Raw.Status.CompletionTime != nil {
			fmt.Printf("job done: %s\n", name)

			return nil
		}

		if j.Raw.Status.Failed > 0 {
			return fmt.Errorf("job fail: %s", name)
		}

		fmt.Printf("job running: %s\n", name)
		time.Sleep(10 * time.Second)
	}
}
