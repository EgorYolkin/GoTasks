FROM golang:1.23.6

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o app .

CMD ["./app"]