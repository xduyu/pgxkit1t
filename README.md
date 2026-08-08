<div align="center">

# pgxkit

**A lightweight Fiber v3 REST API with raw `pgx` access to PostgreSQL — no ORM, no magic.**

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-00ADD8?style=flat&logo=go)](https://go.dev)
[![Fiber](https://img.shields.io/badge/Fiber-v3-00ACD7?style=flat)](https://github.com/gofiber/fiber)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-pgx%20v5-4169E1?style=flat&logo=postgresql&logoColor=white)](https://github.com/jackc/pgx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat)](./LICENSE)
[![Status](https://img.shields.io/badge/Status-WIP-orange.svg?style=flat)]()

</div>

---

## Overview

`pgxkit1t` is a minimal, dependency-light REST API skeleton for user management, built to show how a clean `pgx`-based data layer can look **without** an ORM. Every SQL query is hand-written and explicit — full control, zero abstraction overhead.

It handles the essentials: creating users with hashed passwords, listing them, editing, and deleting by `id` or `uuid`.

> **Heads up:** this is a work-in-progress / learning project. The API surface and package layout may still shift. See [Roadmap](#roadmap) for what's planned.

##  Features

-  **Password hashing** with bcrypt — plaintext passwords never touch the database
-  **UUID-based identity** alongside numeric IDs, generated per user
-  **Raw SQL** via `pgx` — no ORM, no query builder, no hidden N+1 queries
-  **Fiber v3** — fast, expressive routing
-  **Request timing middleware** out of the box
-  **Clean package separation** — handlers, DB helpers, and shared structures each live in their own layer

##  Tech Stack

| Layer | Library |
| HTTP framework | [Fiber v3](https://github.com/gofiber/fiber) |
| PostgreSQL driver | [pgx v5](https://github.com/jackc/pgx) (via `database/sql`) |
| UUID generation | [google/uuid](https://github.com/google/uuid) |
| Password hashing | [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) |
| Language | Go 1.25+ |

##  Project Structure

```
.
├── main.go                      # entry point, Fiber routes
├── db/
│   ├── InitDB_Production.go     # production DB connection (scaffold)
│   ├── CreateTables.go          # users table bootstrap
│   └── helpers/                 # raw SQL: Create / Get / Edit / Delete
│       └── Helpers_Structures/  # response structs for db helpers
├── handlers/
│   ├── CreateUser.go
│   ├── GetUsers.go
│   ├── EditUser.go
│   ├── DeleteUser.go
│   ├── Checkers/                # bcrypt hashing & verification
│   └── Structures/              # request/response structs for handlers
├── GStructures/                 # shared structures (e.g. EditUser payload)
├── middlewares/
│   └── TimeLogger.go            # request duration logging
├── go.mod / go.sum
└── LICENSE                      # MIT
```

##  Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/xduyu/pgxkit1t.git
cd pgxkit1t
```

### 2. Spin up PostgreSQL

```bash
docker run --name pgxkit-db \
  -e POSTGRES_PASSWORD=yourpassword \
  -e POSTGRES_DB=pgxkit \
  -p 5432:5432 -d postgres
```

### 3. Configure the connection

Connect via `pgx` by setting your own connection string (ideally from environment variables) and passing it to `sql.Open`:

```go
connStr := "user=... password=... dbname=... sslmode=disable"
db, err := sql.Open("pgx", connStr)
```

`db/InitDB_Production.go` is a scaffold for the production setup — wiring it up to `os.Getenv` is on the [roadmap](#roadmap).

### 4. Install dependencies & run

```bash
go mod tidy
go run main.go
```

The server starts on `http://ip:3030`, with the API mounted under `/api/v1`. The `users` table is created automatically on startup.

## 📡 API Reference

Base URL: `http://ip:3030/api/v1`

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/users` | List all users |
| `POST` | `/user` | Create a new user |
| `PATCH` | `/user/id` | Update a user's username by `id` |
| `DELETE` | `/user/id` | Delete a user by `id` |
| `DELETE` | `/user/uuid` | Delete a user by `uuid` |

### Examples

**Create a user**
```bash
curl -X POST http://ip:3030/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"username": "ivan", "password": "secret123"}'
```

**List users**
```bash
curl http://ip:3030/api/v1/users
```
Passwords are never returned — only `id`, `uuid`, and `username`.

**Update a username**
```bash
curl -X PATCH http://ip:3030/api/v1/user/id \
  -H "Content-Type: application/json" \
  -d '{"id": 1, "data": {"username": "new_name"}}'
```

**Delete by id**
```bash
curl -X DELETE http://ip:3030/api/v1/user/id \
  -H "Content-Type: application/json" \
  -d '{"id": 1}'
```

**Delete by uuid**
```bash
curl -X DELETE http://ip:3030/api/v1/user/uuid \
  -H "Content-Type: application/json" \
  -d '{"uuid": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"}'
```

## 🗄️ Schema

```sql
CREATE TABLE IF NOT EXISTS users (
    ID SERIAL PRIMARY KEY,
    UUID TEXT NOT NULL,
    USERNAME VARCHAR(100) NOT NULL,
    PASSWORD TEXT NOT NULL
);
```

## 🗺️ Roadmap

- [ ] Environment-based DB configuration (`.env` / `os.Getenv`), including `InitDB_Production`
- [ ] Login endpoint with JWT issuance
- [ ] Input validation (empty username/password currently pass through)
- [ ] Unique constraints / indexes on `uuid` and `username`
- [ ] Test coverage

## 🤝 Contributing

Issues and pull requests are welcome — this project is still taking shape, so feedback on structure and API design is especially appreciated.

## 📄 License

Distributed under the [MIT License](./LICENSE).
