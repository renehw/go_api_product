# Use the official Golang image as the base image
FROM golang:1.23.5

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]