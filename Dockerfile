FROM golang:1.22-alpine as buildenv

WORKDIR /app/
#RUN apk add build-base

COPY go.mod go.mod
COPY go.sum go.sum
COPY .env.dev .env
RUN go mod download

COPY cmd cmd
COPY internal internal

FROM buildenv as check
WORKDIR /app/

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2

RUN golangci-lint run --timeout 5m --verbose
RUN go test ./...

FROM buildenv as builder

RUN go build -o app cmd/server/main.go

FROM golang:1.22-alpine as app
WORKDIR /bin

RUN apk --update --no-cache add curl
COPY --from=builder /app/app /bin/app
COPY --from=builder /app/.env /bin/.env

EXPOSE 9009

CMD ["/bin/app"]

FROM migrate/migrate:latest as migrate

COPY internal/dao/migrations /migrations