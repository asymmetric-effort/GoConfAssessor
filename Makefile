.PHONY: tree
tree:
	@echo 'create build/file_structure.txt'
	@treehash > build/file_structure.txt
	#@tree > build/file_structure.txt

.PHONY: zip
zip: tree
	@echo 'create build/project_files.zip'
	zip -r ./build/project.zip . -x ".git/*" "build/*"

.PHONY: clean
clean: remove-build-dir create-build-dir
	@echo 'clean: ok'

.PHONY: all
all: lint build

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	go fmt ./...
	go vet ./...

.PHONY: create-build-dir
create-build-dir:
	@mkdir -p ./build &> /dev/null || true

.PHONY: remove-build-dir
remove-build-dir:
	@rm -rf ./build &> /dev/null || true

.PHONY: build/manifest-verify
build/manifest-verify: create-build-dir
	@go build -o build/manifest-verify ./cmd/manifest-verify/main.go

.PHONY: build/assess
build/assess: create-build-dir
	@go build -o build/assess ./cmd/assess/main.go

.PHONY: build/analyze
build/analyze: create-build-dir
	@go build -o build/analyze ./cmd/analyze/main.go

.PHONY: build/create-manifest
build/create-manifest: create-build-dir
	@go build -o build/create-manifest ./cmd/create-manifest/main.go


.PHONY: build
build: build/manifest-verify build/create-manifest build/analyze build/assess
	@echo 'build: ok'