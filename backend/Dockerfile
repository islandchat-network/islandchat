FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o island-engine ./cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/island-engine .
EXPOSE 8080
CMD ["./island-engine"]
