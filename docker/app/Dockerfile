FROM golang:latest

COPY ../../src /app
WORKDIR /app
RUN go mod download
RUN go build -o main ./app
EXPOSE 8080

ENTRYPOINT ["./main"]