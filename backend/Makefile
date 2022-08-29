# コマンド一覧
.PHONY: help
help:
	@cat Makefile

.PHONY: install
install:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: gen
gen:
	swagger generate model -f ../swagger.yaml -m generated_swagger -t ./swagger
