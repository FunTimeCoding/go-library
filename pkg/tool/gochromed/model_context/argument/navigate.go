package argument

type Navigate struct {
	Locator string `json:"url"`
	TabID   string `json:"tab_id"`
}
