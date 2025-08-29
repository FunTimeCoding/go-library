package device

import "slices"

func (d *Device) HasTag(s string) bool {
	return slices.Contains(d.Tags, s)
}
