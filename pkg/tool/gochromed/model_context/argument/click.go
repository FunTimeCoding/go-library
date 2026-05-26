package argument

type Click struct {
	UID      string `json:"uid"`
	TabID    string `json:"tab_id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Snapshot *bool  `json:"snapshot"`
}
