package constant

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	CertificateManagerGroup    = "cert-manager.io"
	CertificateManagerVersion  = "v1"
	CertificateResource        = "certificates"
	CertificateRequestResource = "certificaterequests"
)

var (
	CertificateGVR = schema.GroupVersionResource{
		Group:    CertificateManagerGroup,
		Version:  CertificateManagerVersion,
		Resource: CertificateResource,
	}
	CertificateRequestGVR = schema.GroupVersionResource{
		Group:    CertificateManagerGroup,
		Version:  CertificateManagerVersion,
		Resource: CertificateRequestResource,
	}
)
