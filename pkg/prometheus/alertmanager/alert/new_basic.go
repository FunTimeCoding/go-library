package alert

func NewBasic(
	fingerprint string,
	name string,
	severity string,
	summary string,
) *Alert {
	return &Alert{
		Fingerprint: fingerprint,
		Name:        name,
		Severity:    severity,
		Summary:     summary,
	}
}
