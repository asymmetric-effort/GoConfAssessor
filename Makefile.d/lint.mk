
.PHONY: lint
lint: makefile-lint markdown-lint yamllint jsonlint go-lint
	@echo 'lint:ok'

.PHONY: go-lint
go-lint:
	go fmt ./...
	go vet ./...

makefile-lint:
	@echo 'not implemented'

markdown-lint:
	@echo 'not implemented'

yamllint:
	@echo 'not implemented'

jsonlint:
	@echo 'not implemented'