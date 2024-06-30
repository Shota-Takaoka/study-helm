FROM golang:latest AS builder

ENV PORT 8080

WORKDIR /app
COPY ./ /app
RUN go build -o ./test-app1
ENTRYPOINT [ "/app/test-app1" ]
