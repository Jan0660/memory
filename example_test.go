package memory_test

import (
	"fmt"

	"github.com/jan0660/memory"
)

func ExampleTotalMemory() {
	fmt.Printf("Total system memory: %d\n", memory.TotalMemory())
}
func ExampleFreeMemory() {
	fmt.Printf("Free system memory: %d\n", memory.FreeMemory())
}
func ExampleTotalSwap() {
	fmt.Printf("Free system memory: %d\n", memory.TotalSwap())
}
func ExampleFreeSwap() {
	fmt.Printf("Free system memory: %d\n", memory.FreeSwap())
}
func ExampleAvailableMemory() {
	fmt.Printf("Free system memory: %d\n", memory.AvailableMemory())
}
