package convert

type SlimSeverityCount struct {
	Critical    int `json:"critical"`
	Warning     int `json:"warning"`
	Information int `json:"information"`
}
