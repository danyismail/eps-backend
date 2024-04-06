# Start from the Alpine Linux image
FROM golang:alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose port 8080 to the outside world
EXPOSE 1717

# Command to run the executable
CMD ["./myapp"]

#docker run -v /Users/daniismail/Documents/uploads:/app/uploads -v /Users/daniismail/Documents/backend-logs:/app/app.log  -p 1717:1525 -d eps-backend-api