# REST API для управления задачами
REST API на Go для управления задачами с асинхронным логированием действий через канал и поддержкой graceful shutdown.
## Требования
- Go 1.22+

## Клонирование, сборка и запуск

```bash
# Clone
git clone https://github.com/12ilya12/task-mng.git
cd task-mng

# Build
go build -o task-mng cmd/main.go

# Run (Сервер запустится на порту 8000)
go run cmd/main.go