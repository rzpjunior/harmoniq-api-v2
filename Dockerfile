# Builder
FROM golang:1.19.4-alpine3.17 as builder

# Install dependencies
RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy
RUN go mod vendor

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./api

# Final stage: Create a minimal image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api .

# Command to run the executable
CMD ["./api"]
