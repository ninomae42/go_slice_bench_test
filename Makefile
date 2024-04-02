.PHONY: test

test:
	@echo "Running tests..."
	go test -benchmem -bench=. ./...
