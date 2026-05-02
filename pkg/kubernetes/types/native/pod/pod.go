package pod

import "k8s.io/api/core/v1"

type Pod struct {
	Cluster string
	Name    string
	Raw     *v1.Pod
}
