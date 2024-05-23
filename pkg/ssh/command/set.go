package command

func (c *Command) Set(
	k string,
	v string,
) *Command {
	c.Environment[k] = v

	return c
}
