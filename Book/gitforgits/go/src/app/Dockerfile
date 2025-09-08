# Use the official Go runtime as a parent image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy go.mod and go.sum first (better for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your code
COPY . .

# Build the app (creates binary named "app")
RUN go build -o app .

# Expose port 80 to the outside world
EXPOSE 80

# Run the binary
CMD ["./app"]
