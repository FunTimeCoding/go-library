package collector

import "fmt"

func (c *Collector) StringIdentifier(s string) string {
	return fmt.Sprintf("%s-%s", c.Prefix, s)
}
