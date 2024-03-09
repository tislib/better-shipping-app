# better-shipping-app

Demo: https://better-shipping-app-ui.apibrew.io/

## Setup Backend

1. Clone the repository
2. Install the dependencies
3. Run the server

```bash
# Clone the repository
git clone git@github.com:tislib/better-shipping-app.git
cd better-shipping-app

# Install the dependencies
go mod download
go build -o app cmd/server/main.go

# Run the server
/bin/app
```

### Configuring Database

The application uses PostgreSQL as the database. You can configure the database by setting the .enf file.

### Database Migrations

You can run database migration with go migrate tool.

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -path internal/dao/migrations -database "postgres://user:password@localhost:5432/better_shipping_app?sslmode=disable" up
```

see .env.dev

(you need to copy .env.dev to .env and set the values or you can specify environment as ENV=dev)

## Setup Frontend

1. Install the dependencies
2. Run the server

```bash
cd ui

# Install the dependencies
npm install

# Run the server
npm start
```

## Run the tests

```bash
go test ./...
```

## Docker Image Setup

```bash
docker build -t better-shipping-app .
docker run -p 8080:8080 better-shipping-app
```