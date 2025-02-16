GO := go

.PHONY: default
default: help

.PHONY: deps
deps:
	go mod tidy

.PHONY: generate-models
generate-models: deps
	oapi-codegen -package models -generate types -o internal/models/generated.gen.go schemas/openapi.yaml
    
.PHONY: clear-logs
clear-logs:
	rm -r ./*/*.log

.PHONY: lint
lint: deps
	golangci-lint run -v

.PHONY: test
test:
	TEST_MODE=1 go test -v ./...

.PHONY: run
run:
	sudo docker compose up avito-shop-service --build -d

.PHONY: help
help:
	@echo "Доступные команды:"
	@echo "  generate-models  -  Сгенерировать модели из openapi спецификации"
	@echo "  clear-logs  -  Удалить файлы логов"
	@echo "  lint  -  Запустить линтер"
	@echo "  test  -  Запустить тесты"
	@echo "  deps  -  Установка всех зависимостей(напр. для golangci-lint)"
	@echo "  run   -  Запуск проекта"