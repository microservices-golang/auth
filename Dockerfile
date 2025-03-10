# Этап сборки
FROM golang:1.23-alpine AS build

COPY . /github.com/microservices-golang/auth/source/
WORKDIR /github.com/microservices-golang/auth/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/main.go

# Этап запуска
FROM alpine:latest

WORKDIR /root/
COPY --from=build /github.com/microservices-golang/auth/source/bin/crud_server .

CMD ["./crud_server"]