package response

type StorageContent struct {
	Volume string `json:"volume"`
	Format string `json:"format"`
	Size   uint64 `json:"size"`
}
