# Start from the Alpine Linux image
FROM golang:alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the rest of the application code
COPY . .

# Download dependencies
RUN go mod init && go mod tidy

# Build the Go application
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o myapp

# Expose port 1525 to the outside world :)
EXPOSE 1525

# Command to run the executable
CMD ["./myapp"]

#docker run -v /Users/daniismail/Documents/uploads:/app/uploads -v /Users/daniismail/Documents/backend-logs:/app/app.log  -p 1525:1525 -d eps-backend-api