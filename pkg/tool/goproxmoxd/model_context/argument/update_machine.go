package argument

type UpdateMachine struct {
	Identifier  int    `json:"identifier"`
	Node        string `json:"node"`
	Name        string `json:"name"`
	Tags        string `json:"tags"`
	OnBoot      *bool  `json:"onboot"`
	Cores       int    `json:"cores"`
	Memory      int    `json:"memory"`
	Description string `json:"description"`
	Delete      string `json:"delete"`
}
