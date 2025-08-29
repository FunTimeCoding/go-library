package cache

func (c *Cache) Clear() {
	c.DeviceTypes = nil
	c.DeviceRoles = nil
	c.Tags = nil
	c.Tenants = nil
	c.Sites = nil
}
