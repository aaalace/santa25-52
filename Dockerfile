FROM golang:1.23-alpine

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o santa25-52 ./cmd

EXPOSE 8080

CMD ["./santa25-52"]