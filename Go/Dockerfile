# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.12-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
#COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go get github.com/gorilla/mux
RUN go get github.com/streadway/amqp

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Expose port 8080 to the outside world
EXPOSE 8082

# Run the executable
CMD ["go", "run", "main.go"]