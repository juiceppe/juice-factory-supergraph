# Use an official Golang runtime as a parent image
FROM golang:1.20 as builder

# Set the working directory to /go/src/app
WORKDIR /app

# Copy the entire directory contents into the container at /go/src/app
COPY ./subgraphs/ingredients-subgraph/ ./

# Download dependencies
RUN go mod download

# Install the package
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
# Set the environment variable for the PORT
FROM scratch

# Expose port 3009 to the outside world
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3007

ENTRYPOINT [ "/app/main" ]