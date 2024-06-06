FROM golang:1.20

WORKDIR /app
COPY . .

RUN go build -o auth ./cmd/auth/main.go

CMD ["./auth"]
