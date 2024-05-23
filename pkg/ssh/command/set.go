package command

func (c *Command) Set(
	k string,
	v string,
) {
	c.Environment[k] = v
}
