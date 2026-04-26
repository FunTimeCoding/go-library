package response

type EventEntry struct {
	Type    string            `json:"type"`
	Payload EventEntryPayload `json:"data"`
}
