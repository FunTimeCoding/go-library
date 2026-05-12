package argument

import "time"

func (i *Instance) Duration(
	name string,
	value time.Duration,
	usage string,
) {
	i.flags.Duration(name, value, usage)
}
