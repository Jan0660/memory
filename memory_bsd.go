//go:build freebsd || openbsd || dragonfly || netbsd
// +build freebsd openbsd dragonfly netbsd

package memory

func sysTotalMemory() uint64 {
	s, err := sysctlUint64("hw.physmem")
	if err != nil {
		return 0
	}
	return s
}

func sysFreeMemory() uint64 {
	s, err := sysctlUint64("hw.usermem")
	if err != nil {
		return 0
	}
	return s
}

// Not implemented
func sysTotalSwap() uint64 {
	return 0
}

// Not implemented
func sysFreeSwap() uint64 {
	return 0
}

func sysAvailableMemory() uint64 {
	return sysFreeMemory()
}
