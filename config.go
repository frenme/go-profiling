package main

import (
	"flag"
	"time"
)

type Config struct {
	Port          string
	EnableCPUProf bool
	EnableTrace   bool
	CPUProfFile   string
	TraceFile     string
	Timeout       time.Duration
}

func parseFlags() *Config {
	config := &Config{}
	
	flag.StringVar(&config.Port, "port", "8080", "HTTP server port")
	flag.BoolVar(&config.EnableCPUProf, "cpu-prof", true, "Enable CPU profiling")
	flag.BoolVar(&config.EnableTrace, "trace", true, "Enable execution tracing")
	flag.StringVar(&config.CPUProfFile, "cpu-file", "cpu.prof", "CPU profile output file")
	flag.StringVar(&config.TraceFile, "trace-file", "trace.out", "Trace output file")
	flag.DurationVar(&config.Timeout, "timeout", 30*time.Second, "Server timeout")
	
	flag.Parse()
	return config
}