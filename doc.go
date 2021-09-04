// Package memory provides a single method reporting total system memory
// accessible to the kernel.
package memory

// TotalMemory returns the total accessible system memory in bytes.
//
// The total accessible memory is installed physical memory size minus reserved
// areas for the kernel and hardware, if such reservations are reported by
// the operating system.
//
// If accessible memory size could not be determined, then 0 is returned.
func TotalMemory() uint64 {
	return sysTotalMemory()
}

// FreeMemory returns the total free system memory in bytes.
//
// The total free memory is installed physical memory size minus reserved
// areas for other applications running on the same system.
//
// If free memory size could not be determined, then 0 is returned.
func FreeMemory() uint64 {
	return sysFreeMemory()
}

// FreeSwap returns the free amount of swap(page file on Windows) in bytes.
func FreeSwap() uint64 {
	return sysFreeSwap()
}

// TotalSwap returns the total amount of swap(page file on Windows) in bytes.
func TotalSwap() uint64 {
	return sysTotalSwap()
}

// AvailableMemory returns free + buffer/cache memory on Linux, on other platforms it's just free memory.
func AvailableMemory() uint64 {
	return sysAvailableMemory()
}
