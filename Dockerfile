FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stressteste ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/stressteste .
CMD ["./stress_test", "cargateste", "-u", "URL_AQUI", "-r", "NUM_REQUESTS", "-c", "CONCURRENCY"]