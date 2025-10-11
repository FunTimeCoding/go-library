package response

type Statistic struct {
	Summary struct {
		BytesProcessedPerSecond int     `json:"bytesProcessedPerSecond"`
		LinesProcessedPerSecond int     `json:"linesProcessedPerSecond"`
		TotalBytesProcessed     int     `json:"totalBytesProcessed"`
		TotalLinesProcessed     int     `json:"totalLinesProcessed"`
		ExecTime                float64 `json:"execTime"`
		QueueTime               float64 `json:"queueTime"`
		Subqueries              int     `json:"subqueries"`
		TotalEntriesReturned    int     `json:"totalEntriesReturned"`
	} `json:"summary"`
	Querier struct {
		Store struct {
			TotalChunksRef        int `json:"totalChunksRef"`
			TotalChunksDownloaded int `json:"totalChunksDownloaded"`
			ChunksDownloadTime    int `json:"chunksDownloadTime"`
			Chunk                 struct {
				HeadChunkBytes    int `json:"headChunkBytes"`
				HeadChunkLines    int `json:"headChunkLines"`
				DecompressedBytes int `json:"decompressedBytes"`
				DecompressedLines int `json:"decompressedLines"`
				CompressedBytes   int `json:"compressedBytes"`
				TotalDuplicates   int `json:"totalDuplicates"`
			} `json:"chunk"`
		} `json:"store"`
	} `json:"querier"`
	Ingester struct {
		TotalReached       int `json:"totalReached"`
		TotalChunksMatched int `json:"totalChunksMatched"`
		TotalBatches       int `json:"totalBatches"`
		TotalLinesSent     int `json:"totalLinesSent"`
		Store              struct {
			TotalChunksRef        int `json:"totalChunksRef"`
			TotalChunksDownloaded int `json:"totalChunksDownloaded"`
			ChunksDownloadTime    int `json:"chunksDownloadTime"`
			Chunk                 struct {
				HeadChunkBytes    int `json:"headChunkBytes"`
				HeadChunkLines    int `json:"headChunkLines"`
				DecompressedBytes int `json:"decompressedBytes"`
				DecompressedLines int `json:"decompressedLines"`
				CompressedBytes   int `json:"compressedBytes"`
				TotalDuplicates   int `json:"totalDuplicates"`
			} `json:"chunk"`
		} `json:"store"`
	} `json:"ingester"`
}
