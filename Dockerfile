# Use an official Go runtime as a parent image, matching your Go version
FROM golang:1.20.4

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 5000 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/app/main"]
