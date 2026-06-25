package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) StorageContent(
	s *proxmox.Storage,
) ([]*proxmox.StorageContent, error) {
	return s.GetContent(c.context)
}
