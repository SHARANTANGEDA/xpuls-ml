# Use the official Go 1.21.1 image as a base image
FROM golang:1.21.1 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Set GOPROXY environment variable (optional)
# This step is good if you're building behind a corporate firewall or if you want to ensure the Go modules are fetched directly from the Go proxy for faster build times.
ENV GOPROXY=https://proxy.golang.org

# Build the Go app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o xpuls-ml-server .

# Start fresh from a smaller image to create a smaller final container
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the output from the builder stage
COPY --from=builder /app/xpuls-ml-server .
RUN mkdir -p schema
COPY --from=builder /app/schema schema
RUN chmod a+x /root/xpuls-ml-server


# Command to run the executable
ENTRYPOINT ["./xpuls-ml-server", "serve"]
