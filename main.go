package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"runtime/trace"
	"syscall"
	"time"
	_ "net/http/pprof"
)

func main() {
	config := parseFlags()
	
	if config.EnableCPUProf {
		cpuFile, err := os.Create(config.CPUProfFile)
		if err != nil {
			log.Fatal("Failed to create CPU profile file:", err)
		}
		defer cpuFile.Close()
		
		if err := pprof.StartCPUProfile(cpuFile); err != nil {
			log.Fatal("Failed to start CPU profile:", err)
		}
		defer pprof.StopCPUProfile()
	}
	
	if config.EnableTrace {
		traceFile, err := os.Create(config.TraceFile)
		if err != nil {
			log.Fatal("Failed to create trace file:", err)
		}
		defer traceFile.Close()
		
		if err := trace.Start(traceFile); err != nil {
			log.Fatal("Failed to start trace:", err)
		}
		defer trace.Stop()
	}
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/cpu", cpuIntensiveHandler)
	mux.HandleFunc("/io", ioIntensiveHandler)
	mux.HandleFunc("/memory", memoryIntensiveHandler)
	mux.HandleFunc("/fibonacci", fibonacciHandler)
	mux.HandleFunc("/stats", statsHandler)
	mux.HandleFunc("/benchmark", benchmarkHandler)
	
	handler := loggingMiddleware(headerMiddleware(mux))
	
	fmt.Printf("Server starting on :%s\n", config.Port)
	if config.EnableCPUProf {
		fmt.Printf("CPU profiling enabled, output: %s\n", config.CPUProfFile)
	}
	if config.EnableTrace {
		fmt.Printf("Tracing enabled, output: %s\n", config.TraceFile)
	}
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		log.Fatal(http.ListenAndServe(":"+config.Port, handler))
	}()
	
	<-c
	fmt.Println("\nShutting down gracefully...")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Go Profiling Demo</title>
</head>
<body>
    <h1>Go Profiling Demo Server</h1>
    <h2>Available Endpoints:</h2>
    <ul>
        <li><a href="/cpu">/cpu</a> - CPU intensive computation</li>
        <li><a href="/io">/io</a> - I/O intensive operation</li>
        <li><a href="/memory">/memory</a> - Memory allocation test</li>
        <li><a href="/fibonacci">/fibonacci</a> - Fibonacci calculation</li>
        <li><a href="/stats">/stats</a> - System statistics</li>
        <li><a href="/benchmark?workers=5&iterations=1000">/benchmark</a> - Configurable benchmark</li>
    </ul>
    <h2>Profiling Endpoints:</h2>
    <ul>
        <li><a href="/debug/pprof/">/debug/pprof/</a> - Main pprof page</li>
        <li><a href="/debug/pprof/goroutine">/debug/pprof/goroutine</a> - Goroutine profile</li>
        <li><a href="/debug/pprof/heap">/debug/pprof/heap</a> - Heap profile</li>
        <li><a href="/debug/pprof/profile">/debug/pprof/profile</a> - CPU profile</li>
    </ul>
</body>
</html>
`)
}

func cpuIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var result int
	for i := 0; i < 1000000; i++ {
		result += i * i
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "CPU intensive calculation result: %d\nDuration: %v\n", result, duration)
}

func ioIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://httpbin.org/delay/1")
	if err != nil {
		http.Error(w, "Failed to make request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	
	time.Sleep(100 * time.Millisecond)
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}
	
	fmt.Fprintf(w, "I/O operation completed, response size: %d bytes\n", len(body))
}

func memoryIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	var slices [][]byte
	for i := 0; i < 100; i++ {
		data := make([]byte, 1024*1024)
		for j := range data {
			data[j] = byte(i + j)
		}
		slices = append(slices, data)
	}
	
	totalSize := 0
	for _, slice := range slices {
		totalSize += len(slice)
	}
	
	fmt.Fprintf(w, "Memory allocation completed, total allocated: %d bytes\n", totalSize)
}

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	result := fibonacciCalculation(35)
	duration := time.Since(start)
	fmt.Fprintf(w, "Fibonacci(35) = %d\nDuration: %v\n", result, duration)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	printSystemInfo()
	fmt.Fprintf(w, "System statistics logged to console\n")
}