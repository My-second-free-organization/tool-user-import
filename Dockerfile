FROM golang:1.22-alpine AS build
RUN apk add --no-cache git ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/platform-api ./cmd/platform-api

FROM alpine:3.19
RUN addgroup -g 1001 app && adduser -u 1001 -G app -D app
WORKDIR /app
COPY --from=build /bin/platform-api .
USER app
EXPOSE 8080
ENTRYPOINT ["./platform-api"]
