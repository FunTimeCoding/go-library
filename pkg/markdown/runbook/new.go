package runbook

func New(b *[]byte) *Runbook {
	return &Runbook{source: b}
}
