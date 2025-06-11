package profile

type Profile struct {
	Profile string
	Account string `json:"name"`
	Email   string `json:"email"`
}
