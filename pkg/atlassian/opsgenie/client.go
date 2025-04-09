package opsgenie

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/alert_enricher"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/client"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/user_map"
	"github.com/funtimecoding/go-library/pkg/face"
)

type Client struct {
	context context.Context
	webHost string

	teamMap *team_map.Map
	userMap *user_map.Map

	userClient *client.Client
	teamClient *client.Client

	shortUser         face.StringAlias
	shortAlert        face.StringAlias
	descriptionToName face.StringAlias
	tagToTeam         face.SliceAlias
	parseDescription  constant.ParseDescription

	enricher *alert_enricher.Enricher
}
