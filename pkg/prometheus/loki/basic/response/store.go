package response

type Store struct {
	TotalChunksReference  int   `json:"totalChunksRef"`
	TotalChunksDownloaded int   `json:"totalChunksDownloaded"`
	ChunksDownloadTime    int   `json:"chunksDownloadTime"`
	Chunk                 Chunk `json:"chunk"`
}
