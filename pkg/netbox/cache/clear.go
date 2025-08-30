package cache

func (c *Cache) Clear() {
	c.DeviceRoles = nil
	c.DeviceTypes = nil
	c.InternetAddresses = nil
	c.Manufacturers = nil
	c.PhysicalAddresses = nil
	c.Prefixes = nil
	c.Racks = nil
	c.Sites = nil
	c.Tags = nil
	c.Tenants = nil
}
