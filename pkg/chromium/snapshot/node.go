package snapshot

type Node struct {
	UID              string
	Role             string
	Name             string
	Value            string
	BackendDOMNodeID int64
	Children         []*Node
}
