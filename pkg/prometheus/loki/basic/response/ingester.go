package response

type Ingester struct {
	TotalReached       int   `json:"totalReached"`
	TotalChunksMatched int   `json:"totalChunksMatched"`
	TotalBatches       int   `json:"totalBatches"`
	TotalLinesSent     int   `json:"totalLinesSent"`
	Store              Store `json:"store"`
}
