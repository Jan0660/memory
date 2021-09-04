//go:build !linux && !darwin && !windows && !freebsd && !dragonfly && !netbsd && !openbsd
// +build !linux,!darwin,!windows,!freebsd,!dragonfly,!netbsd,!openbsd

package memory

func sysTotalMemory() uint64 {
	return 0
}
func sysFreeMemory() uint64 {
	return 0
}

func sysTotalSwap() uint64 {
	return 0
}

func sysFreeSwap() uint64 {
	return 0
}

func sysAvailableMemory() uint64 {
	return 0
}
