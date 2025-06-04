.PHONY: build/manifest-verify
build/manifest-verify: create-build-dir
	@go build -o build/manifest-verify ./cmd/manifest-verify/main.go