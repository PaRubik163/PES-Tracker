# PES-Tracker

# Описание

PES Tracker — это веб-сервис для учёта личных доходов и расходов, написанный на Go с использованием Gin, GORM, PostgreSQL и Redis.
Сервис предоставляет REST API и веб-интерфейс, позволяя пользователям удобно управлять своими финансами: добавлять, удалять и просматривать транзакции.

Приложение реализует авторизацию на основе JWT-токенов, что обеспечивает безопасность API и изоляцию пользовательских данных.
Каждый пользователь видит только свои операции и баланс.

# Предназначение

Пользователь хочет отслеживать свои финансовые операции и анализировать траты.
PES Tracker помогает:

- добавлять доходы и расходы с категориями и описаниями;

- автоматически считать общий баланс;

- безопасно работать с данными через JWT-аутентификацию.

# Основные технологии

- Go — язык серверной логики

- Gin — фреймворк для построения REST API

- GORM — ORM для работы с PostgreSQL

- PostgreSQL — база данных для хранения пользователей и транзакций

- Redis - база данных для хранения сессий клиентов в кэше

- JWT (JSON Web Token) — механизм аутентификации и авторизации

- HTML / JS / CSS — простая фронтенд-часть

# Архитектура
Проект построен по многоуровневому принципу, что облегчает поддержку и масштабирование:

- Handlers (controllers) — обработка HTTP-запросов, парсинг и валидация данных.

- Service (business logic) — бизнес-логика приложения (расчёт баланса, добавление и удаление доходов и расходов, валидация бизнес-правил).

- Repository (data access) — слой работы с PostgreSQL через GORM и Redis.

- Auth (JWT middleware) — реализация JWT-аутентификации и авторизации пользователей.

# Установка
- Для установки нужно выбрать директорию, где будет проект проекта:  
  ```cd <your_dir>```
- Потом необходимо выполнить эту команду:  
  ```git clone https://github.com/PaRubik163/PES-Tracker.git``` 
- В выбранной папке появится папка ```PES-Tracker``` c проектом.

# Работа с API
# Переменные среды
Сначала необходимо переименовать файл ./.env.example на ./.env и установить параметры:  
- **DB_HOST** - хост PostgreSQL
- **DB_PORT** - порт PostgreSQL
- **DB_USER** - юзер PostgreSQL
- **DB_PASS** - пароль PostgreSQL
- **DB_NAME** - имя базы данных PostgreSQL
- **REDIS_HOST** - хост Redis
- **REDIS_PORT** - порт Redis
- **REDIS_PASS** - пароль Redis
- **GIN_ADDR** - хост, на котором будет работать сервис 
- **JWT_KEY** - секретный ключ для генерации и валидации JWT токена

**Запуск**
- Для запуска API и веб-интерфейса необходимо выбрать директорию проекта:  
```cd <путь к папке PES-Tracker>```
- Далее надо запустить файл ```cmd/main.go```:  
```go run cmd/main.go```

**Управление**  

***Регистрация и вход***  

***Регистрация***  
- Здесь и далее предполагается, что **GIN_ADDR**=localhost:8080

```http
POST http://localhost:8080/api/v1/register
Content-Type: application/json

{
  "login": "<логин>",
  "password": "<пароль>"
}
```
Как работает   

Пользователь отправляет POST запрос на ```/api/v1/register``` с телом:  

```{"login": "<логин>","password": "<пароль>"}```  

В ответ получает 201 Created(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только POST): ``404 page not found``.
- **Неправильное тело запроса**: ``invalid request``.
- **Такой пользователь уже существует**: ``user already exists``.
- **Ошибки с паролем** (не удовлетворяет условиям) ``invalid password``.

***Вход***
```http
POST http://localhost:8080/api/v1/login
Content-Type: application/json

{
  "login": "<логин>",
  "password": "<пароль>"
}
```
Как работает  

Пользователь отправляет POST запрос на ```/api/v1/login``` с телом:  

```{"login": "<логин>","password": "<пароль>"}```  
- Тело ответа:
```http
{
  "message":"Welcome to the PES Tracker",
  "token": "<JWT токен>",
  "user": {
        "create_session_at": "<время создания сессии>",
        "id": "<выданное айди>",
        "login": "<логин>"
  }
}
```

В ответ получает 200 OK и **JWT токен для последующей авторизации**(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только POST): ``404 page not found``.
- **Неправильное тело запроса**: ``invalid request``.
- **Неправильное данные в теле запроса**: ``invalid login or password``.
- **Такой пользователь не существует**: ``user doesn't exists``.

Для всех операций, кроме **login** и **register**, нужен заголовок Authorization:
```Authorization: Bearer <полученный при login токен>```
- ***Токен действует 24 часа***
- Если не передать токен, то возвращает ошибку (401 Unauthorized): ```missing token```
- Если токен неправильный, то возвращает ошибку (401 Unauthorized): ```token is malformed```

***Выход***
```http
POST http://localhost:8080/api/v1/logout
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
```
Как работает  

Пользователь отправляет POST запрос на ```/api/v1/logout``` с заголовком:  
```Authorization: Bearer <полученный при login токен>```
- Тело ответа:
```http
{"message": "logout successful"}
```
В ответ получает 200 OK(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только POST): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Что-то пошло не так**: ```500 error```
  
***Узнать информацию о пользователе***
```http
GET http://localhost:8080/api/v1/me
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
```
Как работает  

Пользователь отправляет GET запрос на ```/api/v1/me``` с заголовком:
```Authorization: Bearer <полученный при login токен>```  
- Тело ответа:
```http
{
    "created_session_at": "<время создания сессии>",
    "expenses_month": "<сумма трат пользователя>",
    "id": "<айди пользователя>",
    "income_month": "<суммарный доход пользователя>",
    "login": "<login>",
    "subscriptions_quantity": "<количество записей о подписках пользователя>",
    "token": "<токен>"
}
```
В ответ получает 200 OK(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только GET): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Истекший токен**: ```not record found on redis```
- **Что-то пошло не так**: ```500 error```

***Добавить подписку***
```http
POST http://localhost:8080/api/v1/new_subscription
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
{
    "name": "<название подписки>", - required
    "amount": "<цена>", - required
    "url": "<url>", - optional
    "start_date": "<дата начала подписки>", - required (2025-08-09T00:00:00Z)
    "billing_period": "<период платежей>", - required (рекомендуется использовать day, month, year)
    "next_billing_date": "<дата следующего платежа>" - required (2025-09-09T00:00:00Z)
}
```
Как работает  

Пользователь отправляет POST запрос на ```/api/v1/new_subscription``` с заголовком:
```Authorization: Bearer <полученный при login токен>```  
- Тело ответа:
```http
{
    "message": "subscription successful created"
}
```
В ответ получает 201 Created(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только POST): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Ошибка в теле запроса**: ```invalid request```
- **Что-то пошло не так**: ```500 error```

***Посмотреть все подписки***
```http
GET http://localhost:8080/api/v1/subscriptions
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
```
Как работает  

Пользователь отправляет GET запрос на ```/api/v1/subscriptions``` с заголовком:
```Authorization: Bearer <полученный при login токен>```  
- Тело ответа:
```http
[
    {
        "id": "<id задачи>",
        "UserID": "<id юзера>",
        "name": "<название подписки>",
        "amount": "<цена подписки>",
        "url": "<ссылка на ресурс, если указывали при создании>",
        "start_date": "<дата начала подписки>",
        "billing_period": "<период платежа>",
        "next_billing_date": "<дата следующего платежа>"
    }
]
```
В ответ получает 200 OK(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только GET): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Ошибка в теле запроса**: ```invalid request```
- **Что-то пошло не так**: ```500 error```

***Добавить запись о доходе***
```http
POST http://localhost:8080/api/v1/new_income
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
{
    "title": "<название дохода>", - required
    "description": "<описание>", - optional
    "amount": "<число прибыли>", - required
    "income_date": "<дата дохода>", - required (2025-09-09T00:00:00Z)
}
```
Как работает  

Пользователь отправляет POST запрос на ```/api/v1/new_income``` с заголовком:
```Authorization: Bearer <полученный при login токен>```  
- Тело ответа:
```http
{
    "message": "income added successfully"
}
```
В ответ получает 201 Created(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только POST): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Ошибка в теле запроса**: ```invalid request```
- **Что-то пошло не так**: ```500 error```

***Посмотреть весь доход***
```http
GET http://localhost:8080/api/v1/income
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
```
Как работает  

Пользователь отправляет GET запрос на ```/api/v1/income``` с заголовком:
```Authorization: Bearer <полученный при login токен>```  
- Тело ответа:
```http
[
    {
        "id": "<id записи>",
        "title": "<название записи>",
        "description": "<описание, если указывали при добавлении>",
        "amount": "<число прибыли>",
        "income_date": "<дата дохода>",
        "UserID": "<id юзера>"
    }
]
```
В ответ получает 200 OK(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только GET): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Ошибка в теле запроса**: ```invalid request```
- **Что-то пошло не так**: ```500 error```

***Добавить запись о трате***
```http
POST http://localhost:8080/api/v1/new_expense
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
{
    "title": "<название траты>", -required
    "description": "<описание траты>", -optional
    "category": "<категория траты>", -required (в веб-версии представлен большой выбор категорий)
    "amount": "<число траты>", - required
    "expense_date": "<дата траты>" - required
}
```
Как работает  

Пользователь отправляет POST запрос на ```/api/v1/new_expense``` с заголовком:
```Authorization: Bearer <полученный при login токен>```  
- Тело ответа:
```http
{
    "message": "expense successful created"
}
```
В ответ получает 201 Created(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только POST): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Ошибка в теле запроса**: ```invalid request```
- **Что-то пошло не так**: ```500 error```

***Посмотреть весе траты***
```http
GET http://localhost:8080/api/v1/expenses
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
```
Как работает  

Пользователь отправляет GET запрос на ```/api/v1/expenses``` с заголовком:
```Authorization: Bearer <полученный при login токен>```  
- Тело ответа:
```http
[
    {
        "id": "<id записи>",
        "title": "<название траты>",
        "description": "<описание, если указывали при создании>",
        "category": "<категория траты>",
        "amount": "<число траты>",
        "expense_date": "<дата траты>",
        "UserID": "<id юзера>"
    }
]
```
В ответ получает 200 OK(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только GET): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Ошибка в теле запроса**: ```invalid request```
- **Что-то пошло не так**: ```500 error```

***Удаление записи о подписке***
```http
DELETE http://localhost:8080/api/v1/subscription/:id
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
```
Как работает  

Пользователь отправляет DELETE запрос на ```/api/v1/subscription/:id``` с заголовком:
```Authorization: Bearer <полученный при login токен>```
- Тело ответа:
```http
{
    "message": "subscription successful deleted"
}
```
В ответ получает 200 OK(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только DELETE): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Непраильный id задачи в запросе** - система не даст удалить чужие записи
- **Что-то пошло не так**: ```500 error```

***Удаление записи о доходе***
```http
DELETE http://localhost:8080/api/v1/income/:id
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
```
Как работает  

Пользователь отправляет DELETE запрос на ```/api/v1/income/:id``` с заголовком:
```Authorization: Bearer <полученный при login токен>```
- Тело ответа:
```http
{
    "message": "income successful deleted"
}
```
В ответ получает 200 OK(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только DELETE): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Непраильный id задачи в запросе** - система не даст удалить чужие записи
- **Что-то пошло не так**: ```500 error```

***Удаление записи о трате***
```http
DELETE http://localhost:8080/api/v1/expense/:id
Content-Type: application/json
Authorization: Bearer <токен, полученный при авторизации>
```
Как работает  

Пользователь отправляет DELETE запрос на ```/api/v1/expense/:id``` с заголовком:
```Authorization: Bearer <полученный при login токен>```
- Тело ответа:
```http
{
    "message": "expense successful deleted"
}
```
В ответ получает 200 OK(в случае успеха), а ошибки в случае:
- **Неправильный метод** (должен быть только DELETE): ```404 page not found```.
- **Пропущенный токен**: ```missing token```
- **Непраильный id задачи в запросе** - система не даст удалить чужие записи
- **Что-то пошло не так**: ```500 error```
