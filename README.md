
# go_final_project

## Описание проекта

**go_final_project** — это веб-приложение для планирования задач, разработанное на языке Go. Приложение позволяет пользователям создавать, просматривать, редактировать и удалять задачи, а также отмечать их как выполненные. В качестве базы данных используется **SQLite** с чисто Go-драйвером [`modernc.org/sqlite`](https://gitlab.com/cznic/sqlite), что упрощает сборку и развёртывание приложения на разных архитектурах. Приложение предоставляет RESTful API и включает фронтенд для удобного взаимодействия.

### Реализованны все задания со звёздочкой, включая

- **Аутентификация:** Реализован механизм аутентификации с использованием JWT-токенов. Доступ к приложению защищён паролем, который задаётся через переменную окружения `TODO_PASSWORD`.
- **Создание Docker-образа:** Разработан `Dockerfile` для сборки Docker-образа приложения, что упрощает его развёртывание и масштабирование. Готовый образ доступен на [Docker Hub](https://hub.docker.com/r/vladimirvereschagin/go_final_project).
- **Кросс-компиляция и поддержка нескольких архитектур:** Благодаря использованию чисто Go-драйвера для SQLite, приложение поддерживает кросс-компиляцию и сборку мультиархитектурных Docker-образов, что позволяет запускать его на различных платформах, включая `linux/amd64` и `linux/arm64`.

## Требования

- **Go** версии **1.22** или выше
- **Git**
- **Docker** (для запуска в контейнере)

## Установка и запуск

### Клонирование репозитория

```bash
git clone https://github.com/VladimirVereshchagin/go_final_project.git
cd go_final_project
```

### Настройка переменных окружения

Создайте файл `.env` в корне проекта со следующим содержимым:

```bash
TODO_PORT=7540
TODO_DBFILE=data/scheduler.db
TODO_PASSWORD=your_password_here
```

- `TODO_PORT` — порт для запуска веб-сервера (по умолчанию 7540).
- `TODO_DBFILE` — имя файла базы данных SQLite.
- `TODO_PASSWORD` — пароль для доступа к приложению. Если оставить пустым, аутентификация не потребуется.

### Установка зависимостей

```bash
go mod download
```

### Инициализация базы данных

Инициализация базы данных не требуется. Приложение автоматически создаст базу данных при первом запуске в директори `data`.

### Сборка приложения

```bash
go build -o app ./cmd
```

### Запуск приложения

```bash
./app
```

### Доступ к приложению

Откройте браузер и перейдите по адресу:

```bash
http://localhost:7540/
```

Если задан пароль (переменная `TODO_PASSWORD` не пустая), вы будете перенаправлены на страницу авторизации:

```bash
http://localhost:7540/login.html
```

Введите установленный пароль для доступа к приложению.

## Быстрый запуск через Docker Hub

Для быстрого развёртывания приложения можно использовать готовый Docker-образ, доступный на [Docker Hub](https://hub.docker.com/r/vladimirvereschagin/go_final_project). Образ поддерживает архитектуры `linux/amd64` и `linux/arm64`, что позволяет запускать его на различных платформах.

### Запуск контейнера

### Важно

Перед запуском контейнера убедитесь, что указали правильный пароль в переменной окружения `TODO_PASSWORD`, или оставьте её пустой, если хотите запустить приложение без пароля. Эти значения должны совпадать с теми, что указаны в файле `.env`. Это необходимо для корректной авторизации в приложении.
Для хранения базы данных используется директория data, которая уже существует в проекте. Используйте команду для запуска контейнера:

```bash
docker run -d \
  -p 7540:7540 \
  --name go_final_project \
  --env TODO_PORT=7540 \
  --env TODO_DBFILE=data/scheduler.db \
  --env TODO_PASSWORD=your_password_here \
  -v $(pwd)/data:/app/data \
  vladimirvereschagin/go_final_project:latest
```

### Пояснения

- `-p 7540:7540` — проброс порта 7540 для доступа к приложению по адресу `http://localhost:7540/`.
- `--env TODO_PORT=7540` — указываем порт приложения.
- `--env TODO_DBFILE=data/scheduler.db` — подключаем файл базы данных.
- `--env TODO_PASSWORD=your_password_here` — задаём пароль для входа в приложение (или оставляем пустым для запуска без пароля).
- `-v $(pwd)/data:/app/data` — монтируем базу данных на хосте для сохранения данных вне контейнера.

### Доступ через браузер

После запуска контейнера приложение будет доступно по адресу:

```bash
http://localhost:7540/
```

### Остановка и удаление контейнера

Чтобы остановить и удалить контейнер, выполните следующие команды:

```bash
docker stop go_final_project
docker rm go_final_project
```

## Запуск тестов

### Перед запуском тестов

Убедитесь, что приложение не запущено или использует другую базу данных, чтобы избежать конфликтов.

### Запуск тестов через скрипт

Тесты используют отдельную тестовую базу данных `test_data/test_scheduler.db`, чтобы избежать конфликтов с основной базой данных приложения.
Используйте скрипт `run-tests.sh` для автоматического запуска тестов.
Скрипт автоматически обрабатывает случаи с установленным паролем и без него.

```bash
./run-tests.sh
```

## Как работает скрипт `run-tests.sh`

- Запускает приложение в фоне с указанным `TODO_PASSWORD`.
- Устанавливает переменную окружения `TODO_DBFILE` в значение `$(pwd)/test_data/test_scheduler.db`.
- Создаёт директорию `test_data`, если она не существует.
- Запускает приложение в фоне, используя тестовую базу данных.
- Получает JWT-токен для авторизации (если пароль установлен) и устанавливает переменную окружения `TOKEN`.
- Запускает тесты, используя установленные переменные окружения.
- Останавливает приложение после завершения тестов.
- Удаляет тестовую базу данных `test_data/test_scheduler.db`.
- Удаляет директорию `test_data`, если она пуста.

### Настройки для тестов

В файле `tests/settings.go` можно настроить параметры:

- `Port`: порт, на котором запускается приложение (по умолчанию 7540).
- `DBFile`: путь к файлу базы данных для тестирования.
- `Token`: JWT-токен для авторизации. Обычно устанавливается автоматически
  скриптом `run-tests.sh`.

## Дополнительная информация

### CI/CD

В проекте настроен GitHub Actions для автоматической сборки и тестирования при пуше в ветки `main` и `new-feature`.
При успешной сборке создаётся мультиархитектурный Docker-образ и отправляется в Docker Hub.

### Pre-commit хуки

Используется `pre-commit` для автоматической проверки кода перед коммитом. Установите хуки командой:

```bash
pre-commit install
```

Для ручной проверки всего кода выполните:

```bash
pre-commit run --all-files --verbose
```

## Структура проекта

- `cmd/` — точка входа в приложение (`main.go`).
- `internal/` — внутренние пакеты приложения:
  - `app/` — настройка маршрутов и обработчиков.
  - `auth/` — аутентификация и работа с JWT.
  - `config/` — загрузка и управление конфигурацией.
  - `models/` — модели данных.
  - `repository/` — работа с базой данных.
  - `services/` — бизнес-логика приложения.
  - `timeutils/` — функции для работы с датами и временем.
- `tests/` — модульные и интеграционные тесты.
- `web/` — файлы фронтенда (HTML, CSS, JavaScript).

## Обратная связь

Если у вас есть вопросы или предложения, пожалуйста, создайте [issue](https://github.com/VladimirVereshchagin/go_final_project/issues) или [pull request](https://github.com/VladimirVereshchagin/go_final_project/pulls) в репозитории проекта.

## Лицензия

Этот проект распространяется под лицензией MIT. Подробности см. в файле [LICENSE](LICENSE).
