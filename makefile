.PHONY: run-test-server build test-j-mode clean

build:
	go build -o longtest

run-test-server:
	@echo "Running test server on :8080..."
	go run test_server/test_server.go

test-j-mode: build
	@echo "Testing J mode..."
	URL='http://localhost:8080' MODE=J ./longtest

clean:
	rm -f longtest

