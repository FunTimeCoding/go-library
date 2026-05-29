package service

type ListQuery struct {
	ResourceType  string
	Namespace     string
	AllNamespaces bool
	LabelSelector string
	FieldSelector string
}
