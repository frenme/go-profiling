.PHONY: build run clean profile trace

BINARY_NAME=profiling-demo
BUILD_DIR=build

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

run-no-profile:
	./$(BUILD_DIR)/$(BINARY_NAME) -cpu-prof=false -trace=false

clean:
	rm -rf $(BUILD_DIR)
	rm -f *.prof *.out

profile:
	go tool pprof $(BUILD_DIR)/$(BINARY_NAME) cpu.prof

trace:
	go tool trace trace.out

deps:
	go mod tidy

test:
	go test -v ./...

install: build
	cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/