package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/constant"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (t *Tester) AddArgocdApp(
	name string,
	namespace string,
	syncStatus string,
	healthStatus string,
) {
	app := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "argoproj.io/v1alpha1",
			"kind":       "Application",
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": constant.Argocd,
			},
			"spec": map[string]interface{}{
				"destination": map[string]interface{}{
					"namespace": namespace,
				},
			},
			"status": map[string]interface{}{
				"sync": map[string]interface{}{
					"status": syncStatus,
				},
				"health": map[string]interface{}{
					"status": healthStatus,
				},
			},
		},
	}
	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group:    "argoproj.io",
			Version:  "v1alpha1",
			Resource: "applications",
		},
	).Namespace(constant.Argocd).Create(
		t.context(),
		app,
		v1.CreateOptions{},
	)
	errors.PanicOnError(e)
}
