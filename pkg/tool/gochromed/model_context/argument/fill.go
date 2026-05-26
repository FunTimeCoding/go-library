package argument

type Fill struct {
	UID      string `json:"uid"`
	Value    string `json:"value"`
	TabID    string `json:"tab_id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Snapshot *bool  `json:"snapshot"`
}
