package service

type ApplyQuery struct {
	Manifest  string
	Namespace string
	Override  bool
	DryRun    bool
}
