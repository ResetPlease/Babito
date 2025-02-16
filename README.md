# 🛍️ Babito - Merch Store Service

**Сервис для покупки фирменного мерча с интеграцией платежей и авторизации**

![CI/CD](https://github.com/ResetPlease/Babito/actions/workflows/go.yml/badge.svg)
![Test Coverage](https://img.shields.io/badge/coverage-67.3%25-green)

## Инструкция по запуску

### Запуск через Make
```bash
make run
```

### Запуск через Docker Compose
```bash
sudo docker compose up avito-shop-service --build -d
```

## 🧪 Тестирование

### Запуск тестовой среды
1. Запустите тестовую БД:
```bash
sudo docker compose up db-test --build -d
```

2. Запустите тесты:
```bash
make test
```

### Особенности тестов
- Расположение: рядом с хендлерами [`api/handlers/*_test.go`](./api/handlers/)
- Покрытые кейсы:
  - ✅ Перевод средств между пользователями
  - ✅ Покупка мерча
  - ✅ Авторизация и аутентификация
  - ✅ Обработка ошибок (Bad Request, недостаточно средств и т.д.)

### Покрытие кода
```bash
TEST_MODE=1 go test -cover ./...
api/handlers    coverage: 67.3% 
api/middleware  coverage: 23.5%
```

## 🔧 Архитектура и инфраструктура

### Ключевые компоненты
| Компонент               | Назначение                          |
|-------------------------|-------------------------------------|
| [PostgreSQL](https://www.postgresql.org/)              | Основное хранилище данных           |
| [PgBouncer](https://www.pgbouncer.org/)               | Пул соединений для БД               |
| [Github Actions](https://github.com/ResetPlease/Babito/actions)| CI/CD пайплайн                      |
| [Golangci-lint](https://golangci-lint.run/)           | Статический анализ кода             |

### Особенности реализации
- 🐘 Оптимизация подключений через [PgBouncer](https://www.pgbouncer.org/) (конфиг в [`docker-compose.yaml`](./docker-compose.yaml))
- 📦 Миграций как таковых нет, инициализация через [`init.sql`](./migrations/init.sql)
- 📊 Провел нагрузочное тестирование с  `Yandex.Tank`([Документация](https://yandextank.readthedocs.io/en/latest/)) и генератором нагрузки `phantom`. **Ответы всех ручек не превышают 30мс. 
С отчетом можно ознакомиться здесь [LOAD_TESTING.md](./load_test/LOAD_TESTING.md)**
- 🔒 Упрощенное хеширование паролей (sha256 вместо bcrypt для демо-целей)

## 📚 Дополнительные материалы
- [Postman коллекция](./postman/Babito.postman_collection.json) - описанные запросы
- [Отчет по нагрузочному тестированию](./load_test/LOAD_TESTING.md) - результаты нагрузочного тестирования

## ❓ FAQ

### Почему нет конфиг-файла?
На текущий момент используется минимальная конфигурация через переменные окружения и код

### Где можно посмотреть статус CI/CD?
На вкладке [Actions](https://github.com/ResetPlease/Babito/actions)
