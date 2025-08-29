package device

func (d *Device) RemoveTag(s string) []string {
	var tags []string

	for _, t := range d.Tags {
		if t != s {
			tags = append(tags, t)
		}
	}

	d.Tags = tags

	return d.Tags
}
