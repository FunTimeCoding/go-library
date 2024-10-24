package keepass

import "github.com/tobischo/gokeepasslib/v3"

type Client struct {
	database *gokeepasslib.Database
}
