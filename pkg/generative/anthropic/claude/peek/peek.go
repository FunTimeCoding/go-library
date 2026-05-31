package peek

type Peek struct {
	LineCount        int
	UserMessageCount int
	Entries          []Entry
	ToolCounts       []ToolCount
	TotalToolCalls   int
}
