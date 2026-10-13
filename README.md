# GoTickts

GoTickts is a REST API for event management and ticket bookings. It supports user registration and login, JWT-protected booking endpoints, and PostgreSQL persistence.

## Features

- Register users and log in to receive a JWT.
- Create, list, retrieve, and update events.
- Book tickets and list the authenticated user's bookings.
- Validate JSON requests and return structured error responses.
- Automatically migrate user, event, and booking tables at startup with GORM.

## Requirements

- Go 1.27 or newer, matching the version in `go.mod`.
- PostgreSQL and a database created for the application.

## Run locally

Create a `.env` file in the project root:

```dotenv
port=8080
dsn="host=localhost user=postgres password=your_password dbname=gotickts port=5432 sslmode=disable TimeZone=UTC"
JWT_SECRET="replace-with-a-long-random-secret"
```

The application reads lowercase `port` and `dsn` values. The database named in the DSN must already exist. Set a private, non-empty `JWT_SECRET` for your environment.

Download dependencies and run the API:

```bash
go mod download
go run ./cmd
```

The server listens at `http://localhost:8080` with the example configuration. On startup, it connects to PostgreSQL and runs GORM AutoMigrate. `GET /` returns a JSON greeting.

## API overview

All routes below are prefixed with `/api/v1`. Requests and responses use JSON. JWT-protected routes require `Authorization: Bearer <token>`.

| Method | Path | Description | Authentication |
| --- | --- | --- | --- |
| `POST` | `/auth/register` | Register a user | Public |
| `POST` | `/auth/login` | Log in and receive a JWT | Public |
| `GET` | `/auth` | List users | Public |
| `GET` | `/auth/me` | Get the current user's profile | JWT required |
| `POST` | `/events` | Create an event | Public |
| `GET` | `/events` | List events | Public |
| `GET` | `/events/:id` | Get an event by ID | Public |
| `PATCH` | `/events/:id` | Update an event | Public |
| `POST` | `/bookings` | Book event tickets | JWT required |
| `GET` | `/bookings/me` | List the current user's bookings | JWT required |

Event creation and update routes are currently public.

### Register and log in

```bash
curl -X POST http://localhost:8080/api/v1/auth/register -H 'Content-Type: application/json' -d '{"name":"Alex Example","email":"alex@example.com","password":"change-me"}'
curl -X POST http://localhost:8080/api/v1/auth/login -H 'Content-Type: application/json' -d '{"email":"alex@example.com","password":"change-me"}'
```

Use the `token` returned by login for protected routes.

### Create an event

`starts_at` uses RFC3339 format; `price` is an integer.

```bash
curl -X POST http://localhost:8080/api/v1/events -H 'Content-Type: application/json' -d '{"title":"Live Music","description":"An evening show","location":"Main Hall","starts_at":"2026-12-01T19:00:00Z","total_tickets":100,"price":25}'
```

### Book tickets

```bash
curl -X POST http://localhost:8080/api/v1/bookings -H 'Content-Type: application/json' -H 'Authorization: Bearer <token>' -d '{"event_id":1,"quantity":2}'
```

Booking requests use the authenticated user from the JWT. The API returns a conflict if there are not enough tickets available.

## Project structure

```text
cmd/main.go                         Application entry point
internal/auth/                      JWT service and claims
internal/config/                    Environment and PostgreSQL configuration
internal/domain/user/               User entity, DTOs, repository, service, handlers, routes
internal/domain/event/              Event entity, DTOs, repository, service, handlers, routes
internal/domain/booking/            Booking entity, DTOs, repository, service, handlers, routes
internal/httpresponse/              Shared error response type
internal/middlewares/               JWT authentication middleware
internal/server/                    Echo setup, validation, migrations, server startup
```

Each domain uses a handler-service-repository structure. Handlers bind and validate HTTP input, services implement application behavior, and repositories use GORM to access PostgreSQL. Domain-specific request and response types live in each domain's `dto` package.

## Data model

- **User:** name, unique email, bcrypt-hashed password.
- **Event:** title, description, location, start time, total and available ticket counts, integer price.
- **Booking:** user and event IDs, quantity, total price, status, unique booking code.

GORM models include IDs and timestamps. The server migrates these models at startup; there is no separate migration command currently.