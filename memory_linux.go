//go:build linux
// +build linux

package memory

import (
	"strconv"
	"syscall"

	"bufio"
	"log"
	"os"
	"strings"
)

func sysTotalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Totalram) * uint64(in.Unit)
}

func sysFreeMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Freeram) * uint64(in.Unit)
}

// Not implemented
func sysTotalSwap() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Totalswap) * uint64(in.Unit)
}

// Not implemented
func sysFreeSwap() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Freeswap) * uint64(in.Unit)
}

func sysAvailableMemory() uint64 {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		// credit: https://github.com/shirou/gopsutil/blob/master/mem/mem_linux.go#L68
		fields := strings.Split(s.Text(), ":")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		value = strings.Replace(value, " kB", "", -1)
		if key == "MemAvailable" {
			memoryAvailable, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return 0
			}
			return memoryAvailable * 1024
		}
	}
	err = s.Err()
	if err != nil {
		return 0
	}
	return 0
}
