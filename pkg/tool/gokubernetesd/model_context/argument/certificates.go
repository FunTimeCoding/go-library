package argument

type Certificates struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Unfiltered bool   `json:"unfiltered"`
}
