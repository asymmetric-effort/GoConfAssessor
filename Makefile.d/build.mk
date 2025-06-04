.PHONY: build
build: build/manifest-verify build/create-manifest build/analyze build/assess
	@echo 'build: ok'