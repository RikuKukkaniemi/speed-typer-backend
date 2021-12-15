# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /go/src/app
COPY . .
RUN go get -d -v
RUN go build -v
CMD ["./speed-typer-backend"]