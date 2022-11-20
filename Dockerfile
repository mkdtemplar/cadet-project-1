# Start from golang base image
FROM golang:latest

# Add Maintainer info
LABEL maintainer="Ivan Markovski"

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . /app


# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...
# Build the Go app
RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build" ]