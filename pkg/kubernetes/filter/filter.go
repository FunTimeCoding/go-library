package filter

type Filter struct {
	Clusters   []string
	Namespaces []string
	Pods       []string
	Containers []string
	Names      []string
}
