# Use an official Node runtime as a parent image
FROM golang:1.22-bullseye

# Set the working directory in the container
WORKDIR /usr/src/app

# Bundle app source
COPY . .

# Setup git inside container
RUN apt update && apt upgrade -y && \
    apt install -y git openssh-client

# Setup bash inside container
RUN apt update && apt install -y bash bash-completion
RUN echo "source /etc/bash/bash_completion.sh" >> ~/.bashprofile
RUN sed -i 's#^root:x:0:0:root:/root:/bin/ash#root:x:0:0:root:/root:/bin/bash#' /etc/passwd