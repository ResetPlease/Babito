GO := go

.PHONY: default
default: help

.PHONY: generate-models
generate-models:
	oapi-codegen -package models -generate types -o internal/models/generated/models.go schemas/openapi.yaml
    
.PHONY: help
help:
	@echo "Доступные команды:"
	@echo "  generate-models  -  Сгенерировать модели из openapi спецификации"
