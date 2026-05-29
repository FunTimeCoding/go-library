package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (t *Tester) AddPodMetrics(
	namespace string,
	name string,
	containers []map[string]string,
) {
	var containerList []interface{}

	for _, c := range containers {
		containerList = append(
			containerList,
			map[string]interface{}{
				"name": c["name"],
				"usage": map[string]interface{}{
					"cpu":    c["cpu"],
					"memory": c["memory"],
				},
			},
		)
	}

	metrics := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "metrics.k8s.io/v1beta1",
			"kind":       "PodMetrics",
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": namespace,
			},
			"containers": containerList,
		},
	}
	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group: "metrics.k8s.io", Version: "v1beta1", Resource: "pods",
		},
	).Namespace(namespace).Create(t.context(), metrics, v1.CreateOptions{})
	errors.PanicOnError(e)
}
