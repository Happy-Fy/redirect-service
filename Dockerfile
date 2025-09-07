FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o redirect ./main.go

EXPOSE 8080

CMD ["./redirect"]