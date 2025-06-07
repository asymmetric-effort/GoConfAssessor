.PHONY: build/assess
build/assess: create-build-dir
	@go build -o build/assess ./cmd/assess/main.go
