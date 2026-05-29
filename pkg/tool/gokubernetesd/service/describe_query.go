package service

type DescribeQuery struct {
	ResourceType string
	Name         string
	Namespace    string
	Unfiltered   bool
}
