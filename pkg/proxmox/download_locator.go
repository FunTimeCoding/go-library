package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) DownloadLocator(
	s *proxmox.Storage,
	content string,
	filename string,
	l string,
) (*proxmox.Task, error) {
	return s.DownloadURL(c.context, content, filename, l)
}
