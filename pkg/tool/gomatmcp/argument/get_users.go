package argument

type GetUsers struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
