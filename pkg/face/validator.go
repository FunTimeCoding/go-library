package face

type Validator interface {
	Validate()
	Concerns() []string
	HasConcerns() bool
	HasConcern(string) bool
}
