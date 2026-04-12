package operation

import (
	"k8s.io/client-go/kubernetes"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

func NetworkPolicy(
	c *kubernetes.Clientset,
	namespace string,
) networkingv1.NetworkPolicyInterface {
	return c.NetworkingV1().NetworkPolicies(namespace)
}
