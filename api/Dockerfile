# Start from golang base image
FROM golang:1.19-alpine

# Add Maintainer info
LABEL maintainer="ShipsRoutes application"

# Setup folders
RUN mkdir cadet-project-v2
WORKDIR cadet-project-v2

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY ./configurations/app.env .


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