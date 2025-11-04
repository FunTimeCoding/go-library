package vault

import "github.com/hashicorp/vault-client-go"

type Client struct {
	client *vault.Client
}
