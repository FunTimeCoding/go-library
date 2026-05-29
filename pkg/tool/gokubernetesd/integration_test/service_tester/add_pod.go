package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/integration_test/service_tester/pod"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
)

func (t *Tester) AddPod(p *pod.Pod) {
	now := time.Now()
	var labelMap map[string]interface{}

	if p.Labels != nil {
		labelMap = make(map[string]interface{})

		for k, v := range p.Labels {
			labelMap[k] = v
		}
	}

	metadata := map[string]interface{}{
		"name":              p.Name,
		"namespace":         p.Namespace,
		"creationTimestamp": now.Format(time.RFC3339),
	}

	if labelMap != nil {
		metadata["labels"] = labelMap
	}

	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Pod",
			"metadata":   metadata,
			"status": map[string]interface{}{
				"phase": p.Phase,
				"containerStatuses": []interface{}{
					map[string]interface{}{
						"name":         p.Name,
						"restartCount": p.Restarts,
						"ready":        true,
						"state": map[string]interface{}{
							"running": map[string]interface{}{
								"startedAt": now.Format(time.RFC3339),
							},
						},
					},
				},
			},
		},
	}
	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group: "", Version: "v1", Resource: "pods",
		},
	).Namespace(p.Namespace).Create(t.context(), u, v1.CreateOptions{})
	errors.PanicOnError(e)
	_, f := t.Clientset.CoreV1().Pods(p.Namespace).Create(
		t.context(),
		&core.Pod{
			ObjectMeta: v1.ObjectMeta{
				Name:      p.Name,
				Namespace: p.Namespace,
				Labels:    p.Labels,
			},
			Status: core.PodStatus{Phase: core.PodPhase(p.Phase)},
		},
		v1.CreateOptions{},
	)
	errors.PanicOnError(f)
}
