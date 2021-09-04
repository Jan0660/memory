package memory_test

import (
	"fmt"
	"testing"

	"github.com/jan0660/memory"
)

func ExampleTotalMemory() {
	fmt.Printf("Total system memory: %d\n", memory.TotalMemory())
}
func ExampleFreeMemory() {
	fmt.Printf("Free system memory: %d\n", memory.FreeMemory())
}
func ExampleTotalSwap() {
	fmt.Printf("Total system swap: %d\n", memory.TotalSwap())
}
func ExampleFreeSwap() {
	fmt.Printf("Free system swap: %d\n", memory.FreeSwap())
}
func ExampleAvailableMemory() {
	fmt.Printf("Available system memory: %d\n", memory.AvailableMemory())
}

func RunExamples(t *testing.T) {
	ExampleAvailableMemory()
	ExampleTotalMemory()
	ExampleFreeMemory()
	ExampleTotalSwap()
	ExampleFreeSwap()
}
