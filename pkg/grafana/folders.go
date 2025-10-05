package grafana

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/grafana/grafana-openapi-client-go/client/folders"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func (c *Client) Folders() []*models.FolderSearchHit {
	result, e := c.client.Folders.GetFolders(&folders.GetFoldersParams{})
	errors.PanicOnError(e)

	return result.Payload
}
