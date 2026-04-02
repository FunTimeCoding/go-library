package response

type EventEntry struct {
	Type string         `json:"type"`
	Data EventEntryData `json:"data"`
}
