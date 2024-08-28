FROM golang:1.23 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY ./server ./server
COPY go.mod .
COPY go.sum .

RUN go mod vendor
RUN go build -mod=vendor -o ./build/server lifelogger/server

# Build a small image
FROM alpine:3.14

# install packages
RUN apk add --no-cache ca-certificates tzdata curl

WORKDIR /app

COPY --from=builder /app/build /app/build
COPY .env .

EXPOSE 22250

ENTRYPOINT ["/bin/sh", "-c" , "./build/server"]
