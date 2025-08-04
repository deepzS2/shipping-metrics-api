FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o /main ./cmd/api

FROM alpine:3.14

WORKDIR /

COPY --from=builder /main /main
COPY migrations ./migrations
EXPOSE 8080
RUN adduser -D api
USER api

ENTRYPOINT [ "/main" ]
