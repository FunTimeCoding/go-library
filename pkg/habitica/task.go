package habitica

type Task struct {
	ID        string          `json:"id"`
	Text      string          `json:"text"`
	Type      string          `json:"type"`
	Notes     string          `json:"notes"`
	Tags      []string        `json:"tags"`
	Completed bool            `json:"completed"`
	Checklist []ChecklistItem `json:"checklist"`
	Streak    int             `json:"streak"`
	Value     float64         `json:"value"`
}
