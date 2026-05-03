package response

type Summary struct {
	BytesProcessedPerSecond int     `json:"bytesProcessedPerSecond"`
	LinesProcessedPerSecond int     `json:"linesProcessedPerSecond"`
	TotalBytesProcessed     int     `json:"totalBytesProcessed"`
	TotalLinesProcessed     int     `json:"totalLinesProcessed"`
	ExecTime                float64 `json:"execTime"`
	QueueTime               float64 `json:"queueTime"`
	Subqueries              int     `json:"subqueries"`
	TotalEntriesReturned    int     `json:"totalEntriesReturned"`
}
