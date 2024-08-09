# syntax=docker/dockerfile:1

#
# build stage
#
FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
ENV CGO_ENABLED=0 \
    GOOS=linux
RUN go build -o /dist/dngl-api cmd/dngl-api/dngl-api.go

#
# final stage
#
FROM scratch
ENV REDIS_DB_URL=redis:6379 \
    DNGL_API_PORT=8080 \
    DNGL_API_HOST=dngl-api
COPY --from=build /dist/dngl-api /
ENTRYPOINT ["/dngl-api"]
