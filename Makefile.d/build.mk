.PHONY: build
build: build/verify-manifest build/create-manifest build/analyze build/assess
	@echo 'build: ok'