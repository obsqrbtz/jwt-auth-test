# JWT Auth Test

## Запуск

Демо API доступно по адресу `https://api.obsqrbtz.space/auth/`. Для тестирования можно воспользоваться `Postman` или другим подобным инструментом.

### 1. Конфигурация

В корне создать файл `.env` со следующими переменными:

```bash
TOKEN_SECRET="c2VjcmV0a2V5"
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="44435777"
POSTGRES_HOST="jwt-auth-db"
POSTGRES_PORT="5432"
POSTGRES_DB="jwt-auth-test"
API_PORT="3030"
SMTP_HOST="smtp.mail.ru"
SMTP_PORT="587"
SMTP_SENDER="dan@obsqrbtz.space"
SMTP_PASSWORD="<пароль от почтового ящика>"
```

Переменной `POSTGRES_HOST` должно быть присвоено значение `jwt-auth-db`. Остальные параметры допустимо изменять.

### 2. Сборка приложения

```bash
docker compose build
```

### 3. Запуск сервера аутентификации и базы данных

```bash
docker compose up
```

## Использование

После запуска приложения `API` будет доступен по адресу `<host>:<API_PORT>/auth`.

### Маршруты

**1. POST `<host>:<API_PORT>/auth/create-tokens`** - создает пользователя в БД и возвращает пару `AccessToken, RefreshToken`.

Параметры:

```bash
{
    "email": "адрес электронной почты пользователя",
}
```

**2. POST `<host>:<API_PORT>/auth/refresh-tokens`** - возвращает новую пару `AccessToken, RefreshToken` и обновляет `RefreshToken` пользователя в БД.

Параметры:

```bash
{
    "refresh_token": "Refresh token в base64"
}
```

**3. GET `<host>:<API_PORT>/auth/get-users`** - возвращает список пользователей из БД если `AccessToken` корректен.

В целях отладки Refresh токены не скрыты в теле запроса.

Если запрос поступает с нового IP, на электронную почту пользователя отправляется предупреждение.
