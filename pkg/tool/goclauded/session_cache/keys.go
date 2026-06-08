package session_cache

func (c *Cache) Keys() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := make([]string, 0, len(c.entries))

	for k := range c.entries {
		result = append(result, k)
	}

	return result
}
