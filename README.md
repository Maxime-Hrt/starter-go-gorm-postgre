# Starter Go Gorm PostgreSQL Fiber

This is a starter project for Go with Gorm, PostgreSQL and Fiber. With this project you can start a new project with a good structure and some functionalities already implemented to create, read, update and delete a user.

## How to run

1. Clone this repository

```bash
git clone git@github.com:Maxime-Hrt/starter-go-gorm-postgresql-fiber.git
```

2. Install the dependencies

```bash
go mod download
```

3. Create a `.env` file at the root of the project and add the following environment variables

```bash
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=your_db_host
DB_PORT=your_db_port
```

4. Run the project

```bash
go run main.go
```

## How to use

### Create a user

```bash
curl -X POST http://localhost:3000/users \
-H "Content-Type: application/json" \
-d '{
  "username": "johndoe",
  "email": "johndoe@example.com",
  "password": "password123"
}'
```