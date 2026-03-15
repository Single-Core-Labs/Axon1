FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o gateway .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gateway .
COPY --from=builder /app/axon.yaml .

EXPOSE 8080
CMD ["./gateway"]
