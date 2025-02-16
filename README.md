# Babito - сервис покупки мерча
Магазин мерча / merch store
## Инструкция по запуску
Запустить с помощью `make` инструкции
```bash
make run
```

или `docker compose`
```bash
sudo docker compose up avito-shop-service --build -d
```

## Тестирование
Для запуска тестов
```bash
sudo docker compose up db-test --build # поднимаем тестовую базу
make test # запускаем тесты
```

### Тесты 
описаны рядом с [хендлерами](./api/handlers/) с суффиксом `_test.go`. 
Проверены все кейсы:
- Передача монет
- Покупка мерча
- Авторизация
- Ошибки типа `bad request`

### Тестовое покрытие
```bash
TEST_MODE=1 go test -cover ./...
	github.com/ResetPlease/Babito/internal/tools		coverage: 0.0% of statements
	github.com/ResetPlease/Babito/api/router		coverage: 0.0% of statements
	github.com/ResetPlease/Babito/cmd		coverage: 0.0% of statements
	github.com/ResetPlease/Babito/internal/db		coverage: 0.0% of statements
	github.com/ResetPlease/Babito/internal/models		coverage: 0.0% of statements
	github.com/ResetPlease/Babito/internal/test_core		coverage: 0.0% of statements
ok  	github.com/ResetPlease/Babito/api/handlers	0.064s	coverage: 67.3% of statements
ok  	github.com/ResetPlease/Babito/api/middleware	0.016s	coverage: 23.5% of statements
```

## Общая информация
- Добавлены [Github Actions](https://github.com/ResetPlease/Babito/actions) конфигурация CI/CD для проверки тестов в пулл реквесте
- `Postman` коллекция для этих ручек описана в [Babito.postman_collection.json](./postman/Babito.postman_collection.json)
- Использовал pg_bouncer для управления пулом подключений и добавил его конфигурацию в [docker_compose.yaml](./docker_compose.yaml), так как это было слабым звеном(выявлено при нагрузочном тестировании)
- Осознанно не стал делать настройку через конфиг, потому что по сути тут два поля для него, а просто создавать лишний конфиг ради конфига не хочется
- Для тестов в [docker_compose.yaml](./docker_compose.yaml) прописана конфигурация для отдельной базы данных: `db-test`
- Не стал писать миграции через go, а сразу инициализировал все таблицы в [init.sql](./migrations/init.sql)
- Добавил конфигурацию `golangci-lint` - [.golangci.yaml](./.golangci.yaml)
- В ближайшее время будут отчеты по нагрузочному тестированию с использованием `Yandex.Tank`([Документация](https://yandextank.readthedocs.io/en/latest/)) и генератором нагрузки `phantom`

## Вопросы/решения
- Хеширование с помощью bcrypt увеличивало время ответа ручки в >5 раз, поменял на sha256 только в рамках этого тестового проекта