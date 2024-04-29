FROM golang:1.22-bullseye

# Set the working directory in the container
WORKDIR /usr/src/app

# Bundle app source
COPY . .
