package response

type Chunk struct {
	HeadChunkBytes    int `json:"headChunkBytes"`
	HeadChunkLines    int `json:"headChunkLines"`
	DecompressedBytes int `json:"decompressedBytes"`
	DecompressedLines int `json:"decompressedLines"`
	CompressedBytes   int `json:"compressedBytes"`
	TotalDuplicates   int `json:"totalDuplicates"`
}
