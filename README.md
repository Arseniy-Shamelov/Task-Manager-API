# Task-Manager-API

Небольшой REST API для ведения задач — что-то вроде бэкенда к туду-листу, только без фронта. Писал в основном для себя, чтобы потренироваться в Go и разобраться с чистой архитектурой. Можно регать пользователей, логиниться по JWT, создавать списки, закидывать в них задачи и отмечать их выполненными.

## Стек

- Go 1.25
- Gin — HTTP-роутер
- PostgreSQL + sqlx (без ORM, обычные SQL-запросы)
- JWT для авторизации (`dgrijalva/jwt-go`)
- Viper для конфига + `godotenv` для секретов
- logrus для логов

## Структура

Разделено на три слоя:

```
cmd/            — точка входа
pkg/handler/    — HTTP-ручки, мидлвары, парсинг запросов
pkg/service/    — бизнес-логика
pkg/repository/ — работа с БД
configs/        — yml-конфиг
schema/         — миграции (up/down)
```

В корне лежат модельки (`todo.go`, `user.go`) и обёртка над `http.Server` с graceful shutdown.

## Запуск

### 1. Поднять Postgres

Проще всего через докер:

```bash
docker run --name=todo-db -p 5436:5432 -d --rm postgres
```

Параметры подключения (имя пользователя, базу, пароль и т.д.) задайте сами — через переменные окружения контейнера или как вам удобнее. Главное, чтобы они совпадали с тем, что прописано в `configs/config.yml`.

### 2. Накатить миграции

Использую [golang-migrate](https://github.com/golang-migrate/migrate):

```bash
migrate -path ./schema -database '<postgres-connection-url>' up
```

### 3. Запустить

```bash
go mod download
go run cmd/main.go
```

По умолчанию крутится на `:1000` (порт тоже можно поменять в конфиге).

## Конфиг

`configs/config.yml`:

```yaml
port: "1000"

db:
  username: "postgres"
  host:     "localhost"
  port:     "5436"
  dbname:   "postgres"
  sslmode:  "disable"
```

## Эндпоинты

### Auth (без токена)

| Метод | Путь             | Что делает                |
|-------|------------------|---------------------------|
| POST  | `/auth/sign-up`  | Регистрация               |
| POST  | `/auth/sign-in`  | Логин, возвращает JWT     |

### Списки (нужен заголовок `Authorization: Bearer <token>`)

| Метод  | Путь              | Что делает                  |
|--------|-------------------|-----------------------------|
| POST   | `/api/lists/`     | Создать список              |
| GET    | `/api/lists/`     | Получить все свои списки    |
| GET    | `/api/lists/:id`  | Получить список по id       |
| PUT    | `/api/lists/:id`  | Обновить                    |
| DELETE | `/api/lists/:id`  | Удалить                     |

### Пункты списков

| Метод  | Путь                         | Что делает                     |
|--------|------------------------------|--------------------------------|
| POST   | `/api/lists/:id/items/`      | Добавить пункт в список        |
| GET    | `/api/lists/:id/items/`      | Все пункты списка              |
| GET    | `/api/items/:id`             | Конкретный пункт               |
| PUT    | `/api/items/:id`             | Обновить (title/description/done) |
| DELETE | `/api/items/:id`             | Удалить                        |

## Примеры

Регистрация:

```bash
curl -X POST http://localhost:1000/auth/sign-up \
  -H "Content-Type: application/json" \
  -d '{"name":"Vasya","username":"vasya","password":"12345"}'
```

Логин:

```bash
curl -X POST http://localhost:1000/auth/sign-in \
  -H "Content-Type: application/json" \
  -d '{"username":"vasya","password":"12345"}'
```

В ответе прилетит `{"token":"..."}`, его потом суём в заголовок:

```bash
curl http://localhost:1000/api/lists/ \
  -H "Authorization: Bearer <token>"
```

Если что-то сломается или будет не так работать — пишите в issues, постараюсь посмотреть.
