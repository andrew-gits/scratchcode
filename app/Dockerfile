# syntax=docker/dockerfile:1.4

FROM node:16-alpine as dev

ARG USER_ID
ARG GROUP_ID
ARG USERNAME
ARG GROUPNAME

# Create user with same gid and uid as person building the image to prevent
# permission issues
RUN getent group $GROUP_ID || addgroup --gid $GROUP_ID user
RUN getent passwd $USER_ID || adduser --disabled-password --gecos '' --uid $USER_ID --gid $GROUP_ID user

# Intstall Golang, Vercel Cli, and Angular Cli
RUN apk add --no-cache bash
RUN apk add chromium
ENV CHROME_BIN=/usr/bin/chromium-browser

RUN mkdir /node
WORKDIR /node
RUN npm install -g @angular/cli@14
RUN npm install -g vercel

COPY --from=golang:1.18-alpine /usr/local/go/ /usr/local/go/
ENV GOPATH /usr/local/go
ENV PATH $GOPATH/bin:$PATH

ENV USER_ID $USER_ID
ENV GROUP_ID $GROUP_ID

# Create and define the application's working directory.
WORKDIR /node/app

# Install the application's dependencies into the node_modules's cache directory.
COPY package.json ./
COPY package-lock.json ./
RUN chown -R ${USER_ID} /node
RUN npm install

RUN chown -R $USER_ID:$GROUP_ID /node

USER root
ENTRYPOINT ["/bin/bash", "entrypoint.sh"]
CMD ["./start-dev-servers.bash"]
