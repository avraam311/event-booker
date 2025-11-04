# Event Booker

Сервис для бронирования мест на события.

## Описание

Event Booker — это REST API сервис на Go для управления событиями и бронированиями мест. Сервис позволяет создавать события, бронировать места, подтверждать бронирования и получать информацию о событиях. Включает автоматическую очистку просроченных бронирований через cron-задачу.

## Функциональность

- Создание событий с указанием названия и количества мест
- Бронирование мест на события
- Подтверждение бронирований
- Получение информации о событии (ID, название, оставшиеся места)
- Автоматическая очистка не подтвержденных бронирований через 1 час

## API Endpoints

Базовый путь: `/event-booker/api`

### Создание события
- **POST** `/events`
- Тело запроса:
  ```json
  {
    "name": "Название события",
    "seats_number": 100
  }
  ```
- Ответ: ID созданного события

### Бронирование места
- **POST** `/events/{id}/book`
- Тело запроса:
  ```json
  {
    "person_name": "Имя человека"
  }
  ```
- Ответ: ID бронирования

### Подтверждение бронирования
- **POST** `/events/confirm/{id}`
- Ответ: Сообщение о подтверждении

### Получение информации о событии
- **GET** `/events/{id}`
- Ответ: Информация о событии (ID, название, оставшиеся места)

## Запуск

### Требования
- Docker и Docker Compose
- Go 1.25.3 (для локального запуска)
- PostgreSQL (через Docker)

### С помощью Docker Compose
1. Создайте файл `.env` с переменными окружения:
   ```
2. Запустите сервисы:
   ```bash
   make up
   ```
   Или для сборки с нуля:
   ```bash
   make buildup
   ```

### Остановка
```bash
make down
```

### Локальный запуск
1. Установите зависимости:
   ```bash
   go mod download
   ```
2. Настройте базу данных и переменные окружения
3. Запустите приложение:
   ```bash
   go run cmd/app/main.go
   ```
4. Запустите cron-сервис:
   ```bash
   go run cmd/cron/main.go
   ```

## Конфигурация

Конфигурация загружается из `config/local.yaml` и переменных окружения.

### config/local.yaml
```yaml
server:
  gin_mode: ""
  port: ":8080"

db:
  max_open_conns: 10
  max_idle_conns: 5
  conn_max_lifetime: 30m

cron:
  spec: "@every 5m"
```

## База данных

Используется PostgreSQL. Схема включает две таблицы:

### event
- `id` (SERIAL PRIMARY KEY)
- `name` (TEXT NOT NULL)
- `seats_number` (INTEGER NOT NULL)
- `seats_number_left` (INTEGER NOT NULL)

### book
- `id` (SERIAL PRIMARY KEY)
- `person_name` (TEXT NOT NULL)
- `book` (VARCHAR(15) NOT NULL) — статус бронирования ("not confirmed" или "confirmed")
- `created_at` (TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP)
- `event_id` (INTEGER NOT NULL REFERENCES event(id) ON DELETE CASCADE)

Индекс на `book.event_id` для оптимизации запросов.

## Миграции

Миграции выполняются автоматически при запуске через сервис `migrator` в Docker Compose. Файлы миграций находятся в папке `migrations/`.

## Cron

Cron-сервис запускает задачу каждые 5 минут для очистки просроченных бронирований (не подтвержденных в течение 1 часа). Увеличивает количество доступных мест при удалении бронирования.

## Зависимости

- github.com/go-playground/validator/v10
- github.com/lib/pq
- github.com/robfig/cron/v3
- github.com/wb-go/wbf

## Структура проекта

```
.
├── cmd/
│   ├── app/
│   │   ├── Dockerfile
│   │   └── main.go          # Основное приложение
│   └── cron/
│       ├── Dockerfile
│       └── main.go          # Cron-сервис
├── config/
│   └── local.yaml           # Конфигурация
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── events/      # Обработчики API для событий
│   │   │   └── response.go  # Вспомогательные функции ответов
│   │   └── server/
│   │       └── server.go    # Настройка сервера и маршрутов
│   ├── infra/
│   │   └── cron/
│   │       └── cron.go      # Инфраструктура для cron
│   ├── middlewares/
│   │   └── cors.go          # CORS middleware
│   ├── models/
│   │   └── models.go        # Модели данных
│   ├── repository/
│   │   └── events/          # Репозиторий для работы с БД
│   └── service/
│       └── events/          # Бизнес-логика
├── migrations/              # Миграции БД
├── docker-compose.yaml      # Docker Compose конфигурация
├── go.mod                   # Go модули
├── go.sum                   # Checksums зависимостей
├── Makefile                 # Команды сборки и запуска
└── README.md                # Этот файл
```

## Линтинг

```bash
make lint
