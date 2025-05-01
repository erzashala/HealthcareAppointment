# 🩺 Healthcare Appointment API (Go + Gin + GORM)

This is a RESTful API built in Go for managing patients and their healthcare appointments. It supports CRUD operations, input validation, Swagger documentation, logging, pagination, and real-world validation like preventing appointments in the past.

---

## 🚀 Features

- Gin for HTTP routing
- GORM for ORM and SQLite for persistence
- Swagger UI documentation
- Input validation with `binding` tags
- Pagination for listing endpoints
- Middleware logging for incoming requests
- Proper HTTP error handling and structured responses
- Future-date validation for appointments

---

## 🛠️ Setup Instructions

### ✅ Requirements

- Go 1.20+
- Docker (optional)
- `swag` CLI: `go install github.com/swaggo/swag/cmd/swag@latest`

### 🔁 Run Locally

```bash
git clone https://github.com/erzashala/HealthcareAppointment.git
cd healthcare-appointment
go mod tidy
swag init -g cmd/main.go
go run cmd/main.go
```

### 🐳 Run with Docker

```bash
docker build -t healthcare-api .
docker run -p 8080:8080 healthcare-api
```

Then visit: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 📚 API Endpoints

### 👤 Patients

| Method | Path              | Description              |
|--------|-------------------|--------------------------|
| POST   | `/patients`       | Create a new patient     |
| GET    | `/patients`       | List all patients        |
| GET    | `/patients/:id`   | Get a patient by ID      |
| PUT    | `/patients/:id`   | Update a patient         |
| DELETE | `/patients/:id`   | Delete a patient         |

### 🗓️ Appointments

| Method | Path                        | Description                              |
|--------|-----------------------------|------------------------------------------|
| POST   | `/appointments`             | Create a new appointment                 |
| GET    | `/appointments`             | List all appointments                    |
| GET    | `/appointments/:id`         | Get an appointment by ID                 |
| PUT    | `/appointments/:id`         | Update an appointment                    |
| DELETE | `/appointments/:id`         | Delete an appointment                    |

---

## 💬 Example Requests and Responses

### ➕ Create a Patient

#### Request
```json
POST /patients
{
  "name": "Erza Shala",
  "email": "erza@example.com"
}
```

#### Response
```json
{
  "id": 1,
  "name": "Erza Shala",
  "email": "erza@example.com"
}
```

---

### ➕ Create an Appointment

#### Request
```json
POST /appointments
{
  "patient_id": 1,
  "date_time": "2025-06-01T15:00:00Z",
  "reason": "Routine checkup"
}
```

#### Response
```json
{
  "id": 1,
  "patient_id": 1,
  "date_time": "2025-06-01T15:00:00Z",
  "reason": "Routine checkup"
}
```

> ⚠️ If the `patient_id` doesn’t exist or the `date_time` is in the **past**, a `400 Bad Request` is returned.

---

### 🔄 Pagination Example

#### Request
```http
GET /patients?limit=10&offset=20
```

#### Response
```json
[
  {
    "id": 21,
    "name": "Erza Shala",
    "email": "erza@example.com"
  },
  ...
]
```

---

### unknown field LeftDelim and RightDelim error?!
When you run:
```bash
swag init -g cmd/main.go
```
You’ll see the generated docs/docs.go file includes this:
```bash
LeftDelim: "{{",
RightDelim: "}}",
```
These lines are totally normal — they’re part of the internal template engine that swaggo uses. 
You can ignore/delete them; they have no effect on how your Swagger UI looks or functions.

---
## 📎 Notes

- Swagger UI is available at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- SQLite file is persisted (unless using in-memory mode)
- `id`, `appointments`, `created_at`, and `updated_at` are auto-managed and hidden from input requests
- Appointments cannot be booked in the past — date-time MUST be in the future
- Pagination is available using `limit` and `offset` query parameters
- Swagger request bodies are pre-filled with example data thanks to `example:` tags in model structs
- No authentication or CORS is included, but middleware can be added easily

---
