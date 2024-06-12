# Builder
FROM golang:1.22.4-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go install github.com/air-verse/air@latest
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./api

CMD ["./api"]
