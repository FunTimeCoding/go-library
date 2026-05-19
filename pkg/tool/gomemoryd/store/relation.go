package store

type Relation struct {
	SourceIdentifier int64  `json:"source_identifier"`
	TargetIdentifier int64  `json:"target_identifier"`
	CreatedAt        string `json:"created_at"`
}
