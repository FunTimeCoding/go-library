package alert_filter

type Filter struct {
	Filter      []string
	Active      *bool
	Silenced    *bool
	Inhibited   *bool
	Unprocessed *bool
	Receiver    string
}
