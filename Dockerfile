FROM golang:1.22.0

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o minitwit.go

# Specify the command to run your application
CMD ["./minitwit"]