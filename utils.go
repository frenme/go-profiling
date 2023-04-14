package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func fibonacciCalculation(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciCalculation(n-1) + fibonacciCalculation(n-2)
}

func generateRandomData(size int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	return data
}

func printSystemInfo() {
	fmt.Printf("CPU cores: %d\n", runtime.NumCPU())
	fmt.Printf("Active goroutines: %d\n", runtime.NumGoroutine())
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %d KB\n", bToKb(m.Alloc))
	fmt.Printf("TotalAlloc = %d KB\n", bToKb(m.TotalAlloc))
	fmt.Printf("Sys = %d KB\n", bToKb(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func bToKb(b uint64) uint64 {
	return b / 1024
}