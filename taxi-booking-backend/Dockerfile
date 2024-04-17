# Start from the official Golang image to build our application.
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch.
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Expose port 80 to the outside world.
EXPOSE 81

# Command to run the executable.
CMD ["./main"]
