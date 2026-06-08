package session_cache

func (c *Cache) Delete(identifier string) {
	c.mutex.Lock()
	delete(c.entries, identifier)
	c.mutex.Unlock()
}
