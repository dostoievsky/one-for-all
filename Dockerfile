FROM golang:1.16

WORKDIR /app

COPY . /app

RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]
