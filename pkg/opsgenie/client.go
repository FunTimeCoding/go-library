package opsgenie

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/opsgenie/client"
	"github.com/funtimecoding/go-library/pkg/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team_map"
	"github.com/funtimecoding/go-library/pkg/opsgenie/user_map"
)

type Client struct {
	context context.Context
	webHost string

	teamMap *team_map.Map
	userMap *user_map.Map

	userClient *client.Client
	teamClient *client.Client

	shortUser         constant.StringAlias
	shortAlert        constant.StringAlias
	descriptionToName constant.StringAlias
	tagsToTeam        constant.SliceAlias
}
