# Use the official Go image with the required version
FROM golang:1.22.7

# Set the working directory
WORKDIR /app

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Expose the gRPC port
EXPOSE 50051

# Build and run the server
RUN go build -o server .
CMD ["./server"]
