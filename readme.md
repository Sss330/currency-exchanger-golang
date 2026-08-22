# Currency Exchange API

REST API сервис для управления справочником валют, обменными курсами и проведения операций конвертации. Проект реализован на языке **Go** с использованием стандартной библиотеки `net/http` и реляционной базы данных **SQLite**.

## Технологический стек

* **Язык программирования:** Go
* **HTTP:** стандартная библиотека `net/http`
* **База данных:** SQLite
* **Архитектура:** MVC / Layered Architecture
* **Формат API:** JSON
* **Формат данных для POST/PATCH:** `application/x-www-form-urlencoded`

## Архитектура проекта

Проект разделен на несколько слоев с четкой ответственностью:

```text
Controllers / Handlers
        ↓
     Services
        ↓
   Repositories
        ↓
      SQLite
```

### Controllers / Handlers

Отвечают за:

* обработку HTTP-запросов;
* валидацию входных данных;
* получение path/query parameters;
* преобразование входных данных;
* формирование HTTP-ответов;
* обработку HTTP-статусов `200`, `201`, `400`, `404`, `409`, `500`.

### Services

Содержат бизнес-логику приложения, включая расчет курса конвертации:

1. прямой курс;
2. обратный курс;
3. кросс-курс.

### Repositories

Отвечают за взаимодействие с SQLite посредством SQL-запросов и подготовленных выражений (`prepared statements`).

Использование параметризованных запросов позволяет избежать SQL-инъекций.

### Models / DTOs

Для разных слоев приложения используются отдельные структуры данных. DTO используются для сериализации и десериализации JSON и HTTP-запросов.

---

## Структура базы данных

### `Currencies`

Хранит информацию о поддерживаемых валютах.

| Колонка    | Тип       | Описание                                     |
| ---------- | --------- | -------------------------------------------- |
| `ID`       | `INTEGER` | Первичный ключ, автоинкремент                |
| `Code`     | `VARCHAR` | Уникальный код валюты, например `USD`, `EUR` |
| `FullName` | `VARCHAR` | Полное название валюты                       |
| `Sign`     | `VARCHAR` | Символ валюты, например `$`, `€`, `₽`        |

Индексы:

* `PRIMARY KEY (ID)`
* `UNIQUE INDEX` по `Code`

### `ExchangeRates`

Хранит обменные курсы между валютными парами.

| Колонка            | Тип          | Описание                              |
| ------------------ | ------------ | ------------------------------------- |
| `ID`               | `INTEGER`    | Первичный ключ, автоинкремент         |
| `BaseCurrencyId`   | `INTEGER`    | FK на `Currencies.ID`, базовая валюта |
| `TargetCurrencyId` | `INTEGER`    | FK на `Currencies.ID`, целевая валюта |
| `Rate`             | `DECIMAL(6)` | Курс базовой валюты к целевой         |

Индексы:

* `PRIMARY KEY (ID)`
* `UNIQUE INDEX` по `(BaseCurrencyId, TargetCurrencyId)`

---

# API

## Валюты

### `GET /currencies`

Получить список всех доступных валют.

**Response — `200 OK`:**

```json
[
  {
    "id": 1,
    "name": "United States dollar",
    "code": "USD",
    "sign": "$"
  },
  {
    "id": 2,
    "name": "Euro",
    "code": "EUR",
    "sign": "€"
  }
]
```

---

### `GET /currency/{code}`

Получить информацию о валюте по ее коду.

**Пример:**

```http
GET /currency/EUR
```

**Response — `200 OK`:**

```json
{
  "id": 2,
  "name": "Euro",
  "code": "EUR",
  "sign": "€"
}
```

**Возможные ошибки:**

* `400 Bad Request` — код валюты не указан;
* `404 Not Found` — валюта не найдена.

---

### `POST /currencies`

Добавить новую валюту.

**Content-Type:**

```text
application/x-www-form-urlencoded
```

**Request body:**

```text
name=United States dollar&code=USD&sign=$
```

**Response — `201 Created`:**

```json
{
  "id": 1,
  "name": "United States dollar",
  "code": "USD",
  "sign": "$"
}
```

**Возможные ошибки:**

* `400 Bad Request` — обязательные поля не заполнены;
* `409 Conflict` — валюта с таким кодом уже существует.

---

# Обменные курсы

### `GET /exchangeRates`

Получить список всех установленных обменных курсов.

Ответ содержит информацию о базовой и целевой валюте.

**Response — `200 OK`:**

```json
[
  {
    "id": 1,
    "baseCurrency": {
      "id": 1,
      "name": "United States dollar",
      "code": "USD",
      "sign": "$"
    },
    "targetCurrency": {
      "id": 2,
      "name": "Euro",
      "code": "EUR",
      "sign": "€"
    },
    "rate": 0.92
  }
]
```

---

### `GET /exchangeRate/{pair}`

Получить курс конкретной валютной пары.

**Пример:**

```http
GET /exchangeRate/USDEUR
```

**Возможные ошибки:**

* `400 Bad Request` — некорректный формат валютной пары;
* `404 Not Found` — обменный курс для пары не найден.

---

### `POST /exchangeRates`

Добавить новый обменный курс.

**Content-Type:**

```text
application/x-www-form-urlencoded
```

**Request body:**

```text
baseCurrencyCode=USD&targetCurrencyCode=EUR&rate=0.92
```

**Response — `201 Created`:**

```json
{
  "id": 1,
  "baseCurrency": {
    "id": 1,
    "name": "United States dollar",
    "code": "USD",
    "sign": "$"
  },
  "targetCurrency": {
    "id": 2,
    "name": "Euro",
    "code": "EUR",
    "sign": "€"
  },
  "rate": 0.92
}
```

**Возможные ошибки:**

* `400 Bad Request` — некорректные входные данные;
* `404 Not Found` — одна из указанных валют не существует;
* `409 Conflict` — обменный курс для данной пары уже существует.

---

### `PATCH /exchangeRate/{pair}`

Обновить существующий обменный курс.

**Пример:**

```http
PATCH /exchangeRate/USDEUR
```

**Content-Type:**

```text
application/x-www-form-urlencoded
```

**Request body:**

```text
rate=0.93
```

**Response — `200 OK`:**

```json
{
  "id": 1,
  "baseCurrency": {
    "id": 1,
    "name": "United States dollar",
    "code": "USD",
    "sign": "$"
  },
  "targetCurrency": {
    "id": 2,
    "name": "Euro",
    "code": "EUR",
    "sign": "€"
  },
  "rate": 0.93
}
```

---

# Конвертация валют

### `GET /exchange`

Рассчитать конвертацию заданной суммы из одной валюты в другую.

**Пример:**

```http
GET /exchange?from=USD&to=EUR&amount=100
```

### Алгоритм поиска курса

Сервис использует три сценария.

#### 1. Прямой курс

Если в базе существует:

```text
USD → EUR = 0.92
```

то используется непосредственно этот курс.

```text
100 USD × 0.92 = 92 EUR
```

#### 2. Обратный курс

Если прямого курса нет, сервис проверяет обратную пару:

```text
EUR → USD = 1.087
```

В этом случае курс рассчитывается как:

```text
USD → EUR = 1 / 1.087
```

#### 3. Кросс-курс

Если прямой и обратный курсы отсутствуют, сервис может использовать общую базовую валюту.

Например:

```text
USD → RUB = 90
USD → EUR = 0.92
```

Тогда курс:

```text
RUB → EUR = 0.92 / 90
```

Полученный курс используется для конвертации исходной суммы.

### Response — `200 OK`

```json
{
  "baseCurrency": {
    "id": 1,
    "name": "United States dollar",
    "code": "USD",
    "sign": "$"
  },
  "targetCurrency": {
    "id": 2,
    "name": "Euro",
    "code": "EUR",
    "sign": "€"
  },
  "rate": 0.92,
  "amount": 100.00,
  "convertedAmount": 92.00
}
```

---

# Формат ошибок

Все ошибки API возвращаются в едином JSON-формате:

```json
{
  "message": "Описание ошибки"
}
```

Например:

```http
HTTP/1.1 404 Not Found
Content-Type: application/json
```

```json
{
  "message": "Currency with code EUR not found"
}
```

---

# Запуск проекта

## Требования

Для запуска необходимы:

* Go `1.20+`;
* SQLite3;
* Git.

## Клонирование репозитория

```bash
git clone https://github.com/username/currency-exchange.git
cd currency-exchange
```

## Запуск

```bash
go run cmd/main.go
```

После запуска API будет доступен на настроенном HTTP-порту.

---

# Сборка

Для создания исполняемого файла:

```bash
go build -o currency-exchange ./cmd
```

После этого приложение можно запустить:

```bash
./currency-exchange
```

---

# Запуск на Linux-сервере

Для запуска приложения в фоне:

```bash
nohup ./currency-exchange > app.log 2>&1 &
```

Проверить процесс:

```bash
ps aux | grep currency-exchange
```

Посмотреть логи:

```bash
tail -f app.log
```

После запуска сервис будет доступен по адресу:

```text
http://<server_ip>:<port>
```

---

# Пример использования

Добавим валюты:

```bash
curl -X POST http://localhost:8080/currencies \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "name=United States dollar&code=USD&sign=$"
```

```bash
curl -X POST http://localhost:8080/currencies \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "name=Euro&code=EUR&sign=€"
```

Добавим обменный курс:

```bash
curl -X POST http://localhost:8080/exchangeRates \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "baseCurrencyCode=USD&targetCurrencyCode=EUR&rate=0.92"
```

Выполним конвертацию:

```bash
curl "http://localhost:8080/exchange?from=USD&to=EUR&amount=100"
```

Результат:

```json
{
  "baseCurrency": {
    "id": 1,
    "name": "United States dollar",
    "code": "USD",
    "sign": "$"
  },
  "targetCurrency": {
    "id": 2,
    "name": "Euro",
    "code": "EUR",
    "sign": "€"
  },
  "rate": 0.92,
  "amount": 100,
  "convertedAmount": 92
}
```

---

# Основные особенности

* REST API без сторонних HTTP-фреймворков;
* стандартная библиотека Go `net/http`;
* SQLite для хранения данных;
* разделение приложения на Controller, Service и Repository;
* подготовленные SQL-запросы;
* единый формат ошибок;
* поддержка CRUD-операций со справочником валют;
* управление обменными курсами;
* прямой, обратный и кросс-курс;
* JSON API;
* поддержка `application/x-www-form-urlencoded` для операций создания и обновления.
