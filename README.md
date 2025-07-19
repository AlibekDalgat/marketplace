# Профильное задание на стажировку бэкенд-разработчиком ВК

<!-- ToC start -->
# Описание задачи
Написать приложение, реализующее REST-API условного маркетплейса:
- авторизация пользователя
- регистрация пользователя
- размещения нового объявления
- отображение ленты объявлений

# Реализация
- Применение фреймворка [gin-gonic/gin](https://github.com/gin-gonic/gin).
- Применение СУБД Postgres посредствов библиотеки [sqlx](https://github.com/jmoiron/sqlx) и написанием SQL запросов.
- Контейнеризация с помощью Docker и docker-compose

**Структура проекта:**
```
.
├── internal
│   ├── app         // точка запуска приложения
│   ├── config      // общие конфигурации приложения
│   ├── delivery    // слой обработки запросов
│   ├── models      // структуры сущностей приложения
│   ├── service     // слой бизнес-логики
│   └── repository  // слой взаимодействия с БД
├── pkg
│   └── validation  // методы валидации данных
├── cmd             // точка входа в приложение
└── db              // SQL файлы миграции
```

# Endpoints
- POST /auth/sign-up - регистрация пользователя
    - Тело запроса:
      - login - логин
      - password - пароль
- POST /auth/sign-in - авторизация пользователя
    - Тело запроса:
      - login - логин
      - password - пароль
- POST /api/adverts - размещение объявления
    - Тело запроса:
      - title - заголовок
      - text - текст объявления
      - image - адрес изображения
      - price - цена
- GET /api/adverts - отображение ленты объявлений
    - Параметры запроса:
      - sort - тип сортировки
      - direction - направление сортировки
      - limit - количестно объявлений на страницу
      - page - номер страницы
      - pricemin - минимальная цена
      - pricemax - максимальная цена
     
# Запуск
```
make build && make run
```
После собри и запуска, необходимо применить миграции к базе данных:
```
make migrate-up
```

# Примеры
Запросы сгенерированы командой curl
### 1. POST /auth/sign-up/
**Запрос:**
```
curl --location --request POST 'localhost:8000/auth/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "alibek",
    "password": "Pass1!"
}'
```
**Тело ответа:**
```
{
    "login": "alibek",
    "password": "$2a$10$7fe5e.6EZIjUVIh27zv.vOQZTyqBqzdKWQep74MSx4GLullW.sUtO"
}
```
В ответе возвращается логин и хэш пароля
### 2. POST /auth/sign-in/
**Запрос:**
```
curl --location --request POST 'localhost:8000/auth/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "alibek",
    "password": "Pass1!"
}'
```
**Тело ответа:**
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwIjoxNzUyOTg2OTk4LCJpYXQiOjE3NTI5NDM3OTh9LCJsb2dpbiI6ImFsaWJlayJ9.hAaLZ7FJ5juQYxq5rYrLlbPRzc4FVHccxH2X7ybvhgU"
}
```
Полученный токен далее необходимо
### 3. POST /api/adverts/
**Запрос:**
```
curl --location --request POST 'localhost:8000/api/adverts' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwIjoxNzUyOTg2OTk4LCJpYXQiOjE3NTI5NDM3OTh9LCJsb2dpbiI6ImFsaWJlayJ9.hAaLZ7FJ5juQYxq5rYrLlbPRzc4FVHccxH2X7ybvhgU' \
--data-raw '{
    "title": "uyuy",
    "text": "text",
    "price": 6
}'
```
**Тело ответа:**
```
{
    "title": "uyuy",
    "text": "text",
    "image": null,
    "price": 6,
    "posting_date": "2025-07-19T16:50:45.994062Z",
    "owner": "alibek"
}
```
### 4. GET /api/adverts/
**Запрос:**
```
curl --location --request GET 'localhost:8000/api/adverts?sort=price&direction=asc' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwIjoxNzUyOTg2OTk4LCJpYXQiOjE3NTI5NDM3OTh9LCJsb2dpbiI6ImFsaWJlayJ9.hAaLZ7FJ5juQYxq5rYrLlbPRzc4FVHccxH2X7ybvhgU' \
--data-raw '{
    "title": "uyuy",
    "text": "text",
    "price": 6
}'
```
**Тело ответа:**
```
{
    "data": [
        {
            "title": "uyuy",
            "text": "text",
            "image": null,
            "price": 6,
            "owner": "alibek",
            "isOwner": true
        }
    ]
}
```
