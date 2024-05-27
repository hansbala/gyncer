# Start from the official Golang base image
FROM golang:1.22.2 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
# COPY go.mod go.sum ./
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app using the syncbridge.go file
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./syncbridge.go

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]