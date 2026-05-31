package argument

type CreateContainerSnapshot struct {
	Identifier int    `json:"identifier"`
	Node string `json:"node"`
	Name string `json:"name"`
}
