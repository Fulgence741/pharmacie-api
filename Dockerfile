FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/pharmacie-api .

FROM alpine:latest

RUN apk add --no-cache ca-certificates \
    && addgroup -S app \
    && adduser -S -G app app

WORKDIR /app
COPY --from=builder /out/pharmacie-api .
RUN mkdir -p logs && chown -R app:app /app

USER app

EXPOSE 8080
ENTRYPOINT ["./pharmacie-api"]
