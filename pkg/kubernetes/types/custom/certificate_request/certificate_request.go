package certificate_request

type CertificateRequest struct {
	Name      string
	Namespace string
	Ready     bool
	Message   string
}
