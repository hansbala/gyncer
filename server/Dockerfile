FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

# Copy the configuration file for Air
COPY .air.toml .air.toml

# Copy project config
COPY config/config.toml config/config.toml

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

# # Start from the official Golang base image
# FROM golang:1.22.2 as builder

# # Set the current working directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # Copy the source from the current directory to the working Directory inside the container
# COPY . .

# # Build the Go app using the server.go file
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./

# # Start a new stage from scratch
# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# # Install Air for live reloading
# RUN apk add --no-cache curl
# RUN curl -fLo /usr/bin/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
# RUN chmod +x /usr/bin/air

# WORKDIR /root/

# # Copy the configuration file for Air
# COPY .air.toml .air.toml

# COPY config/config.toml config/config.toml

# # Copy the Pre-built binary file from the previous stage
# COPY --from=builder /app/server .

# # Expose port 8080 to the outside world
# EXPOSE 8080

# # Command to run the executable
# CMD ["air"]