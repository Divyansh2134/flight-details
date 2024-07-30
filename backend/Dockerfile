# Build stage
FROM golang:1.22-bullseye AS goBuild

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Final stage
FROM debian:bullseye-slim

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=goBuild /app/main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
