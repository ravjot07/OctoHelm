# Build Stage
FROM golang:1.18 AS builder

WORKDIR /app
COPY go.mod main.go main_test.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o service-b .

# Final Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/service-b .
EXPOSE 8081
CMD ["./service-b"]
