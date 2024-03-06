# Build Go API
FROM golang:1.18 as builder

# Set Current Working Directory inside the container
WORKDIR /app

# Copy mod and sum files
COPY go.mod go.sum ./

# Download Dependencies
RUN go mod download

# Copy source from current to Working Directory
COPY . .
COPY .env /app/
COPY .env-test /app/


# Build Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./api/api.go

# Setup runtime environment
FROM alpine:latest

WORKDIR /root/

# Copy pre-built binary file
COPY --from=builder /app/api .
COPY --from=builder /app/.env .
COPY --from=builder /app/.env-test .

# Expose port 8080
EXPOSE 8080
ENV OPTION_VALUE=false
# Run the excutable
CMD ["./api","-test","$OPTION_VALUE"]
