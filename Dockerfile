FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chart-streaming ./cmd/chart-streaming


FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/chart-streaming .
COPY migrations ./migrations


EXPOSE 50051
ENTRYPOINT ["./chart-streaming"]
