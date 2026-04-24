package collector

import "fmt"

func (c *Collector) Integer64Identifier(i int64) string {
	return fmt.Sprintf("%s-%d", c.Prefix, i)
}
