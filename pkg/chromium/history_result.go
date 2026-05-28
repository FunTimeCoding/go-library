package chromium

type HistoryResult struct {
	CurrentIndex int64
	Entries      []*HistoryEntry
}
