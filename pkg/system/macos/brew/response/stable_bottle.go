package response

type StableBottle struct {
	Rebuild     int         `json:"rebuild"`
	RootLocator string      `json:"root_url"`
	Files       BottleFiles `json:"files"`
}
