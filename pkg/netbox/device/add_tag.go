package device

func (d *Device) AddTag(s string) []string {
	d.Tags = append(d.Tags, s)

	return d.Tags
}
