# Stage 1
FROM golang:1.22-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

# Stage 2
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server /app/
CMD ["/app/server"]
