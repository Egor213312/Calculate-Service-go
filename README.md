# Calculate-Service-Go

## Описание

**Calculate-Service-Go** — это HTTP-сервис для выполнения арифметических вычислений. Он принимает математические выражения в формате строки, обрабатывает их и возвращает результат. Проект написан на языке Go и предназначен для интеграции с другими сервисами, приложениями или использования в качестве самостоятельного API.

---

## Основные возможности

- **Обработка арифметических выражений**: поддержка операций сложения, вычитания, умножения и деления.
- **JSON API**: простой интерфейс взаимодействия через POST-запросы.
- **Обработка ошибок**: корректная обработка некорректных выражений, деления на ноль и других возможных проблем.
- **Тестирование**: проект содержит готовые тесты для проверки корректности работы.

---

## Инструкция по установке

### Требования
- Go версии 1.18 и выше.
- Установленный [Git](https://git-scm.com/).
- Доступ в интернет для скачивания зависимостей.

### Установка

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/Egor213312/Calculate-Service-go.git
   cd Calculate-Service-go

# 2. Установите зависимости:
go mod tidy

# Инструкция по запуску
Запуск в режиме разработчика
Запустите сервер:

go run cmd/calc_service/main.go
Сервер будет доступен по адресу http://localhost:8080/api/v1/calculate

# Запуск в режиме пользователя
1. Соберите бинарный файл:

go build -o calc_service cmd/calc_service/main.go

2. Запустите его:

./calc_service

# API
# POST /calculate
Принимает математическое выражение и возвращает результат.

# Пример запроса
{
  "expression": "3 + 2 * (1 + 4)"
}

# Пример ответа
{
  "result": 13
}

# Ошибки
Если выражение некорректно:
HTTP 400 Internal Server Error
{
  "error": "Internal server error"
}
Структура проекта:

/calc_service
    /cmd
        /calc_service
            main.go         # Основной файл для запуска сервера
    /internal
        /calculator
            calculator.go   # Логика обработки выражений
    /pkg
        /utils
            utils.go        # Вспомогательные функции
    README.md
    go.mod                # Go модуль

# Тестирование
Для запуска тестов используйте:

go test -v ./...
Тесты проверяют корректность вычислений, обработку ошибок и работу API.

# Пример кода
Вот пример, как отправить запрос к API через curl:

curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{"expression": "5 * (2 + 3)"}'

Ответ:
{
  "result": 25
}

# Ссылки:
Репозиторий на GitHub
Документация по Go
