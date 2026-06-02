package convert

import "github.com/prometheus/alertmanager/api/v2/models"

func Status(v *models.AlertmanagerStatus, includeConfiguration bool) *SlimStatus {
	result := &SlimStatus{}

	if v.Cluster != nil {
		if v.Cluster.Status != nil {
			result.ClusterStatus = *v.Cluster.Status
		}

		result.PeerCount = len(v.Cluster.Peers)
	}

	if v.VersionInfo != nil && v.VersionInfo.Version != nil {
		result.Version = *v.VersionInfo.Version
	}

	if includeConfiguration && v.Config != nil && v.Config.Original != nil {
		result.Configuration = *v.Config.Original
	}

	return result
}
