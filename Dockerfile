FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY docs/ ./docs/
COPY internal/ ./internal/

RUN go build -o /service ./cmd/storage-api/main.go

EXPOSE 8080

CMD ["/service"]