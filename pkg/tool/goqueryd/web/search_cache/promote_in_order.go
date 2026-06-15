package search_cache

func (c *Cache) promoteInOrder(key string) {
	c.removeFromOrder(key)
	c.order = append(c.order, key)
}
