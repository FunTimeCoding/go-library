package system

type DiskUsageReport struct {
	Path  string
	Total uint64
	Free  uint64
	Used  uint64
}
