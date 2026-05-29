package service

type GetQuery struct {
	ResourceType string
	Name         string
	Namespace    string
	Unfiltered   bool
}
