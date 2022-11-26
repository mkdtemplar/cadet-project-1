# To create container:
# Before use of the command in terminal "docker build --tag cadet-project"
# in go.mod file change the name of the module to "cadet-project"
# otherwise it will not build the image.
# After the image is build in terminal run the following command: "docker run --publish 8080:8080 cadet-project" to start the container

# Changing the module name in go.mod will be not required when the branch is merged to master.

# Start from golang base image
FROM golang:1.19-alpine

# Add Maintainer info
LABEL maintainer="Ports application"

# Setup folders
RUN mkdir /cadet-project
WORKDIR /cadet-project

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