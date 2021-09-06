# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.17.0
# ENV TZ="Asia/Ho_Chi_Minh"
# Add Maintainer Info
# LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /me-english

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o bin/main main.go
RUN ls -a
RUN rm /etc/localtime
RUN ln -s /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
# Expose port 8080 to the outside world
EXPOSE 4040

# Command to run the executable
CMD ["./bin/main"]