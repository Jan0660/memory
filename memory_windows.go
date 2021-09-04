//go:build windows
// +build windows

package memory

import (
	"errors"
	"syscall"
	"unsafe"
)

// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-memorystatusex
type memStatusEx struct {
	dwLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func getMemoryStatus() (*memStatusEx, error) {
	kernel32, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return nil, errors.New("Failed to load kernel32.dll")
	}
	// GetPhysicallyInstalledSystemMemory is simpler, but broken on
	// older versions of windows (and uses this under the hood anyway).
	globalMemoryStatusEx, err := kernel32.FindProc("GlobalMemoryStatusEx")
	if err != nil {
		return nil, errors.New("Couldn't find procedure")
	}
	msx := &memStatusEx{
		dwLength: 64,
	}
	r, _, _ := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(msx)))
	if r == 0 {
		return nil, errors.New("Failed to call Windows API")
	}
	return msx, nil
}

func sysTotalMemory() uint64 {
	msx, err := getMemoryStatus()
	if err != nil {
		return 0
	}
	return msx.ullTotalPhys
}

func sysFreeMemory() uint64 {
	msx, err := getMemoryStatus()
	if err != nil {
		return 0
	}
	return msx.ullAvailPhys
}

func sysTotalSwap() uint64 {
	msx, err := getMemoryStatus()
	if err != nil {
		return 0
	}
	return msx.ullTotalPageFile
}

func sysFreeSwap() uint64 {
	msx, err := getMemoryStatus()
	if err != nil {
		return 0
	}
	return msx.ullAvailPageFile
}

func sysAvailableMemory() uint64 {
	return sysFreeMemory()
}
