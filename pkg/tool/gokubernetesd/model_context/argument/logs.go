package argument

type Logs struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Container  string `json:"container"`
	Tail       int    `json:"tail"`
	Since      string `json:"since"`
	Previous   bool   `json:"previous"`
	Timestamps bool   `json:"timestamps"`
	All        bool   `json:"all"`
}
