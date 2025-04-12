package pod

import core "k8s.io/api/core/v1"

type Pod struct {
	Cluster string
	Name    string
	Raw     *core.Pod
}
