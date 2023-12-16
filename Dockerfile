# Use a lightweight version of Go.
FROM golang:1.16-alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go module files and download dependencies.
COPY go.* ./
RUN go mod download

# Copy the source code from the current directory to the /app directory in the container.
COPY . .

# Build the application. Change the directory to where main.go is located.
RUN cd cmd && go build -o /golang-interview

# Expose the port that the server listens on.
EXPOSE 8080

# Run the binary.
CMD ["/golang-interview"]
