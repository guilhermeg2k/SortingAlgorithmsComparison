package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	totalMemKb := bToKb(m.TotalAlloc)
	totalMemMB := bToMb(m.TotalAlloc)
	fmt.Println(totalMemKb, " Kb\n", totalMemMB, " Mb")
}

func bToKb(b uint64) uint64 {
	return b / 1024
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func generateVector(size int, random, invert bool) []int {
	var vector []int
	rand.Seed(time.Now().UnixNano())

	if random {
		for i := size; i > 0; i-- {
			vector = append(vector, rand.Intn(size))
		}
	} else {
		for i := 0; i < size; i++ {
			vector = append(vector, i)
		}
	}

	if invert {
		for i := 0; i < len(vector)/2; i++ {
			j := len(vector) - i - 1
			vector[i], vector[j] = vector[j], vector[i]
		}
	}
	return vector
}

/*
func main() {
	// Print our starting memory usage (should be around 0mb)
	start := time.Now()
	PrintMemUsage()

	var overall [][]int
	for i := 0; i < 4; i++ {

		// Allocate memory using make() and append to overall (so it doesn't get
		// garbage collected). This is to create an ever increasing memory usage
		// which we can track. We're just using []int as an example.
		a := make([]int, 0, 999999)
		overall = append(overall, a)

		// Print our memory usage at each interval
		PrintMemUsage()
		time.Sleep(time.Second)

	}
	elapsed := time.Since(start)
	fmt.Println("Demorou ", elapsed)
	// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
	overall = nil
	PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}*/
