# Todo API Service

REST API для управления задачами и пользователями на Go с интерактивной Swagger-документацией.

## Возможности

- CRUD для задач: создание, чтение, обновление, удаление
- CRUD для пользователей
- Частичное обновление через указатели (`*string`, `*bool`)
- Хранение в памяти (без БД)
- Потокобезопасная работа с пользователями (`sync.RWMutex`)
- Интерактивная документация Swagger UI
- Обёртка успешных ответов через `SuccessResponse`

## Стек

| Технология | Версия | Назначение |
|-----------|--------|-----------|
| Go | 1.22+ | Язык |
| go-chi/chi | v5.0.10 | HTTP-роутер |
| swaggo/swag | 1.16+ | Генератор Swagger-документации |
| swaggo/http-swagger | v2.0.2 | Swagger UI для chi |

## Установка и запуск

### 1. Клонировать репозиторий

```bash
git clone <repo-url>
cd ai-assist-it
```

### 2. Установить зависимости

```bash
go mod download
```

### 3. Установить `swag` CLI (если ещё не установлен)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$HOME/go/bin
```

### 4. Сгенерировать Swagger-документацию

```bash
swag init -g main.go -o docs
```

### 5. Запустить сервер

```bash
go run main.go
```

Сервер запустится на `http://localhost:8080`.

**Swagger UI:** `http://localhost:8080/swagger/`

## Эндпоинты

Все URL имеют префикс `/api/v1`.

### Задачи (`/todos`)

| Метод | Путь | Описание |
|-------|------|----------|
| `GET` | `/api/v1/todos` | Список всех задач |
| `POST` | `/api/v1/todos` | Создать задачу |
| `GET` | `/api/v1/todos/{id}` | Получить задачу по ID |
| `PUT` | `/api/v1/todos/{id}` | Частично обновить задачу |
| `DELETE` | `/api/v1/todos/{id}` | Удалить задачу |

### Пользователи (`/users`)

| Метод | Путь | Описание |
|-------|------|----------|
| `GET` | `/api/v1/users` | Список всех пользователей |
| `POST` | `/api/v1/users` | Создать пользователя |
| `GET` | `/api/v1/users/{id}` | Получить пользователя по ID |
| `PUT` | `/api/v1/users/{id}` | Частично обновить пользователя |
| `DELETE` | `/api/v1/users/{id}` | Удалить пользователя |

## Примеры запросов

### Создать задачу

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "title": "Купить молоко", "description": "Не забыть про скидку"}'
```

**Ответ (201):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Купить молоко",
    "description": "Не забыть про скидку",
    "done": false,
    "created_at": "2025-01-01T12:00:00Z",
    "updated_at": "2025-01-01T12:00:00Z"
  }
}
```

### Список задач

```bash
curl http://localhost:8080/api/v1/todos
```

**Ответ (200):**

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "title": "Купить молоко",
      "description": "Не забыть про скидку",
      "done": false,
      "created_at": "2025-01-01T12:00:00Z",
      "updated_at": "2025-01-01T12:00:00Z"
    }
  ]
}
```

### Частичное обновление задачи

```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"done": true}'
```

**Ответ (200):** возвращается обновлённая задача (title и description сохраняются, меняется только `done`).

### Создать пользователя

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email": "ivan@example.com", "name": "Иван Иванов"}'
```

**Ответ (201):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "ivan@example.com",
    "name": "Иван Иванов",
    "created_at": "2025-01-01T12:00:00Z"
  }
}
```

### Удалить задачу

```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

**Ответ:** `204 No Content`.

## Формат ответов

**Успешные ответы** обёрнуты в `SuccessResponse`:

```json
{
  "success": true,
  "data": { ... }
}
```

**Ошибки** возвращаются как plain text с соответствующим HTTP-кодом:

| Код | Ситуация | Тело |
|-----|----------|------|
| `400 Bad Request` | Невалидный JSON или нечисловой ID | `Неверный ID` / `Невалидный запрос` |
| `404 Not Found` | Задача или пользователь не найдены | `Задача не найдена` / `Пользователь не найден` |

## Структура проекта

```
ai-assist-it/
├── main.go                       # Точка входа, роуты, Swagger UI
├── go.mod                        # Зависимости
├── go.sum                        # Контрольные суммы
├── README.md                     # Этот файл
├── docs/                         # Сгенерированная Swagger-документация
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── internal/
    ├── handlers/
    │   ├── todo.go               # CRUD-обработчики задач с аннотациями
    │   └── user.go               # CRUD-обработчики пользователей с аннотациями
    └── models/
        ├── todo.go               # Модели Todo, CreateTodoRequest, UpdateTodoRequest
        └── user.go               # Модели User, CreateUserRequest, UpdateUserRequest, SuccessResponse, ErrorResponse
```

### Чек-лист структуры проекта

- [x] `main.go` — точка входа, роуты, Swagger UI
- [x] `go.mod`, `go.sum` — управление зависимостями
- [x] `internal/handlers/todo.go` — CRUD-обработчики задач с аннотациями
- [x] `internal/handlers/user.go` — CRUD-обработчики пользователей с аннотациями
- [x] `internal/models/todo.go` — модели задач с `@Description` и `example`
- [x] `internal/models/user.go` — модели пользователей, `SuccessResponse`, `ErrorResponse`
- [x] `docs/docs.go` — сгенерированный Go-код документации
- [x] `docs/swagger.json` — Swagger-спецификация в JSON
- [x] `docs/swagger.yaml` — Swagger-спецификация в YAML
- [x] `README.md` — этот файл

## Обновление Swagger-документации

После изменения аннотаций в handlers или моделях перегенерируйте документацию:

```bash
swag init -g main.go -o docs
```

**Что означают флаги:**

- `-g main.go` — файл с глобальными аннотациями (`@title`, `@version`, `@host`, `@BasePath`).
- `-o docs` — папка для сохранения сгенерированных файлов.

Если после обновления CLI возникают ошибки компиляции в `docs/docs.go` — синхронизируйте версии CLI и библиотеки:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/swag@latest
go mod tidy
swag init -g main.go -o docs
```

## Ограничения

- Данные хранятся **в памяти** процесса — после перезапуска всё пропадает.
- Нет аутентификации и авторизации.
- Нет пагинации и сортировки.
- Порядок элементов в списках не определён (map в Go не гарантирует порядок).
- Ошибки возвращаются как plain text — модель `ErrorResponse` присутствует, но не используется в текущей реализации.
- Валидационные теги `validate:"..."` в моделях носят декларативный характер — валидатор не подключён.

## Лицензия

MIT — см. файл [LICENSE](LICENSE), если добавлен.

## Контакты

- Автор: Sukhanov
- Email: sukhanovov@gmail.com