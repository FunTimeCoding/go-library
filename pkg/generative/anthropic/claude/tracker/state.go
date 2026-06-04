package tracker

type State struct {
	Offset           int64
	Lines            int
	FirstTimestamp   string
	LastTimestamp    string
	Slug             string
	WorkDirectory    string
	Branch           string
	UserMessageCount int
	FirstMessage     string
}
