FROM golang:latest

ARG USER_ID=1000
ARG GROUP_ID=1000
RUN groupadd -g ${GROUP_ID} hostgroup \
    && useradd -u ${USER_ID} -g hostgroup -m -s /bin/bash hostuser

COPY ../../src /app
WORKDIR /app
RUN go mod download
RUN go build -o main ./app
EXPOSE 8080

USER hostuser

ENTRYPOINT ["./main"]