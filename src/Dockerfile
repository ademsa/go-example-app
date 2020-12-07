FROM golang:latest

WORKDIR /app

COPY . .

RUN go build

EXPOSE 3000

CMD ["./go-example-app"]