.PHONY: build/analyze
build/analyze: create-build-dir
	@go build -o build/analyze ./cmd/analyze/main.go
