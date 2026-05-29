package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
)

func (t *Tester) AddCertificate(
	namespace string,
	name string,
	dnsNames []string,
	notAfter time.Time,
	ready bool,
) {
	readyStatus := "False"

	if ready {
		readyStatus = "True"
	}

	var dnsNamesRaw []interface{}

	for _, d := range dnsNames {
		dnsNamesRaw = append(dnsNamesRaw, d)
	}

	cert := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Certificate",
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"dnsNames": dnsNamesRaw,
			},
			"status": map[string]interface{}{
				"notAfter": notAfter.Format(time.RFC3339),
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": readyStatus,
					},
				},
			},
		},
	}
	_, e := t.Dynamic.Resource(
		schema.GroupVersionResource{
			Group: "cert-manager.io", Version: "v1", Resource: "certificates",
		},
	).Namespace(namespace).Create(
		t.context(),
		cert,
		v1.CreateOptions{},
	)
	errors.PanicOnError(e)
}
