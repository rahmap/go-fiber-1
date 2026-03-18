# Fiber REST API

Simple REST API example using Go, Fiber, GORM, MySQL, and Swagger UI.

## Requirements

- Go 1.20+
- MySQL

## Setup

1. Copy [.env.example](I:\Project\Golang\fiber-rest-api\.env.example) to `.env`
2. Adjust the environment values:

```env
APP_NAME=fiber-rest-api
APP_HOST=
APP_PORT=5008
DATABASE_DSN=root:@tcp(127.0.0.1:3306)/your_database_name?parseTime=true
API_TOKEN=replace-with-your-secret-token
```

## Run App

Make sure your MySQL database is running and the database in `DATABASE_DSN` already exists.

Run the application with:

```bash
go run .
```

If the app starts successfully, it will listen on:

```text
http://localhost:5008
```

If you change `APP_PORT`, adjust the URL accordingly.

## Swagger UI

After the app is running, open Swagger UI in your browser:

```text
http://localhost:5008/swagger/index.html
```

Swagger documents these endpoints:

- `GET /api/v1/users`
- `GET /api/v1/users/{id}`
- `POST /api/v1/users`

All user endpoints require the `X-TOKEN` header.

Example header:

```text
X-TOKEN: replace-with-your-secret-token
```

## Example Requests

Get all users:

```bash
curl -X GET "http://localhost:5008/api/v1/users" \
  -H "X-TOKEN: replace-with-your-secret-token"
```

Get user by ID:

```bash
curl -X GET "http://localhost:5008/api/v1/users/1" \
  -H "X-TOKEN: replace-with-your-secret-token"
```

Create user:

```bash
curl -X POST "http://localhost:5008/api/v1/users" \
  -H "Content-Type: application/json" \
  -H "X-TOKEN: replace-with-your-secret-token" \
  -d "{\"name\":\"Rahma AP\",\"age\":24,\"address\":\"Yogyakarta\"}"
```
