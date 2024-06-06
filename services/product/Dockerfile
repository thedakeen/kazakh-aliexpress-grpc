FROM golang:1.20

WORKDIR /app
COPY . .

RUN go build -o product ./cmd/product/main.go

CMD ["./product"]
