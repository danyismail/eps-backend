# Start from the Alpine Linux image
FROM golang:alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose port 8080 to the outside world
EXPOSE 1717

# Command to run the executable
CMD ["./myapp"]

#docker run -v /home/eps/go/src/assembly/eps-backend/uploads:/app/uploads -v /Users/daniismail/Documents/backend-logs:/app/app.log  -p 1525:1525 -d eps-backend-api