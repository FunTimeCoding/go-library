//go:build !windows

package system

import "syscall"

func DiskUsage(path string) *DiskUsageReport {
	// TODO: Make OS-independent: https://github.com/ricochet2200/go-disk-usage
	var s syscall.Statfs_t
	result := &DiskUsageReport{Path: path}

	if e := syscall.Statfs(path, &s); e != nil {
		return result
	}

	result.Total = s.Blocks * uint64(s.Bsize)
	result.Free = s.Bavail * uint64(s.Bsize)
	result.Used = result.Total - (s.Bfree * uint64(s.Bsize))

	return result
}
