# Build Stage
FROM golang:1.18 AS builder

WORKDIR /app
COPY go.mod main.go main_test.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o service-a .

# Final Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/service-a .
EXPOSE 8080
CMD ["./service-a"]