.PHONY: build/manifest-verify
build/verify-manifest: create-build-dir
	@go build -o build/verify-manifest ./cmd/verify-manifest/main.go