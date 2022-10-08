# syntax=docker/dockerfile:1.4

FROM --platform=$BUILDPLATFORM node:16 as dev

ARG USER_ID
ARG GROUP_ID
ARG USERNAME
ARG GROUPNAME
ARG GO_ARCHIVE=go1.19.1.linux-amd64.tar.gz
ARG ARCHIVE_SHA256=acc512fbab4f716a8f97a8b3fbaa9ddd39606a28be6c2515ef7c6c6311acffde

# Create a user and group with the same IDs as the host so files can be shared
COPY create-group-user.bash ./
RUN ./create-group-user.bash

# Intstall Golang, Vercel Cli, and Angular Cli
RUN mkdir /node
WORKDIR /node
RUN npm install -g @angular/cli@14
RUN npm install -g vercel
RUN curl -OL https://go.dev/dl/${GO_ARCHIVE}
RUN tar -C /usr/local -xzf ${GO_ARCHIVE}
RUN [ $(sha256sum ${GO_ARCHIVE} | cut -d' ' -f1) = "${ARCHIVE_SHA256}" ]
RUN rm ${GO_ARCHIVE}

ENV PATH=$PATH:/usr/local/go/bin

# Create and define the application's working directory.
WORKDIR /node/app

# Install the application's dependencies into the node_modules's cache directory.
COPY package.json ./
COPY package-lock.json ./
RUN chown -R ${USER_ID} /node
USER ${USER_ID}
RUN npm install

USER root
ENTRYPOINT ["/bin/bash", "entrypoint.sh"]
CMD ["./start-dev-servers.bash"]

FROM dev as dev-envs

RUN <<EOF
apt-get update
apt-get install -y --no-install-recommends git
EOF

RUN <<EOF
useradd -s /bin/bash -m vscode
groupadd docker
usermod -aG docker vscode
EOF

CMD ./start-dev-servers.bash
