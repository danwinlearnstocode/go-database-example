FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-database-example

EXPOSE 8080

CMD ["/go-database-example"]