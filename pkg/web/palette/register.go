package palette

func (r *Registry) Register(commands ...Command) {
	r.commands = append(r.commands, commands...)
}
