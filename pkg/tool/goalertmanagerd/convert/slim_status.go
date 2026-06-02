package convert

type SlimStatus struct {
	ClusterStatus string `json:"cluster_status"`
	PeerCount     int    `json:"peer_count"`
	Version       string `json:"version"`
	Configuration string `json:"configuration,omitempty"`
}
