FROM golang:1.24-alpine3.20 AS build
WORKDIR /app
COPY src src
COPY go.mod go.mod
COPY go.sum go.sum
RUN go build -o server ./src


FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/server ./server
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT "./server" "--host" "0.0.0.0" "--port" "${PORT}"
