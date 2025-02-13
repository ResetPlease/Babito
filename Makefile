GO := go

.PHONY: default
default: help

.PHONY: generate-models
generate-models:
	oapi-codegen -package models -generate types -o internal/models/generated.go schemas/openapi.yaml
    
.PHONY: clear-logs
clear-logs:
	rm -r ./*/*.log

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: test
test:
	TEST_MODE=1 go test -v ./...

.PHONY: help
help:
	@echo "Доступные команды:"
	@echo "  generate-models  -  Сгенерировать модели из openapi спецификации"
	@echo "  clear-logs  -  Удалить файлы логов"
	@echo "  lint  -  Запустить линтер"
	@echo "  test  -  Запустить тесты"
