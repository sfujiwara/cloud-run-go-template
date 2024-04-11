FROM golang:1.12.9-alpine3.10 AS build

WORKDIR /app

COPY go.mod go.mod
COPY main.go main.go

RUN go build -o server


FROM alpine:3.10.2

WORKDIR /app

COPY --from=build /app/server server
EXPOSE 8080

CMD ["./server"]
