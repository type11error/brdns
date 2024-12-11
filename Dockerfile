# syntax=docker/dockerfile:1
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN make

CMD ["./brdns"]


# Use a minimal image for runtime
FROM debian:bullseye-slim

# Set up a non-root user
RUN useradd -ms /bin/bash brdns
USER brdns

# Set the working directory
WORKDIR /home/brdns

# Copy the built binary from the builder
COPY --from=builder /app/brdns .

# Expose the DNS server port
EXPOSE 8053

# Command to run the application
CMD ["./brdns"]

