package collector

import "fmt"

func (c *Collector) IntegerIdentifier(i int) string {
	return fmt.Sprintf("%s-%d", c.Prefix, i)
}
