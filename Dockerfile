FROM golang:1.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o minitwit ./src/main.go
RUN chmod +x minitwit

# Use a Docker multi-stage build to create a small final image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/minitwit .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./minitwit"]
