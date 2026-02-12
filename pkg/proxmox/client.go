package proxmox

import (
	"context"
	"github.com/luthermonson/go-proxmox"
)

type Client struct {
	context    context.Context
	client     *proxmox.Client
	user       string
	password   string
	token      string
	secret     string
	selfSigned bool
	log        bool
	verbose    bool
}
