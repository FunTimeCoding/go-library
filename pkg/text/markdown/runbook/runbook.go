package runbook

type Runbook struct {
	source *[]byte

	Filename string
	Title    string
	Sections []Section
}
