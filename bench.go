package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func benchmarkHandler(w http.ResponseWriter, r *http.Request) {
	workersStr := r.URL.Query().Get("workers")
	workers := 10
	if workersStr != "" {
		if w, err := strconv.Atoi(workersStr); err == nil {
			workers = w
		}
	}
	
	iterationsStr := r.URL.Query().Get("iterations")
	iterations := 1000
	if iterationsStr != "" {
		if i, err := strconv.Atoi(iterationsStr); err == nil {
			iterations = i
		}
	}
	
	start := time.Now()
	
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var sum int
			for j := 0; j < iterations; j++ {
				sum += j * j
			}
		}()
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Fprintf(w, "Benchmark completed: %d workers, %d iterations each\n", workers, iterations)
	fmt.Fprintf(w, "Total duration: %v\n", duration)
}