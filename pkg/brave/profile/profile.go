package profile

type Profile struct {
	Profile string
	Path    string
	Account string `json:"name"`
	Email   string `json:"email"`
}
