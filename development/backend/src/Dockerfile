FROM golang:1.24

WORKDIR /app

COPY go.mod . 
COPY go.sum . 
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
