FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /app

# Copy all project files into the container
COPY . .

# Download Go module dependencies
RUN go mod tidy

# Build the binary from the main.go file in the cmd folder
RUN go build -o auth-service main.go

# Expose port 50052 for the API Gateway
EXPOSE 50052

# Command to run the built binary
CMD ["./auth-service"]
