# syntax=docker/dockerfile:1.4

FROM --platform=$BUILDPLATFORM node:16 as dev

ARG USER_ID
ARG GROUP_ID
ARG USERNAME
ARG GROUPNAME

COPY create-group-user.bash ./
RUN ./create-group-user.bash

ARG GO_ARCHIVE=go1.19.1.linux-amd64.tar.gz
ARG ARCHIVE_SHA256=acc512fbab4f716a8f97a8b3fbaa9ddd39606a28be6c2515ef7c6c6311acffde

RUN mkdir /node
WORKDIR /node
RUN npm install -g @angular/cli@14
RUN npm install -g vercel
RUN curl -OL https://go.dev/dl/${GO_ARCHIVE}
RUN tar -C /usr/local -xzf ${GO_ARCHIVE}
RUN [ $(sha256sum ${GO_ARCHIVE} | cut -d' ' -f1) = "${ARCHIVE_SHA256}" ]
RUN rm ${GO_ARCHIVE}

ENV PATH=$PATH:/usr/local/go/bin

# Create and define the node_modules's cache directory.
RUN mkdir /cache
WORKDIR /cache
COPY package.json ./
COPY package-lock.json ./

# Install the application's dependencies into the node_modules's cache directory.
RUN chown -R ${USER_ID} /node /cache
USER ${USER_ID}
RUN npm install

# Create and define the application's working directory.
RUN mkdir /node/app
WORKDIR /node/app

USER root
ENTRYPOINT ["/bin/bash", "entrypoint.sh"]
CMD ["/node/app/start-dev-servers.bash"]

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

# example https://medium.com/bb-tutorials-and-thoughts/angular-local-development-with-docker-compose-13719b998e43

# # install Docker tools (cli, buildx, compose)
# COPY --from=gloursdocker/docker / /

# # https://docs.docker.com/config/containers/multi-service_container/
# # https://www.docker.com/blog/how-to-setup-your-local-node-js-development-environment-using-docker/
# # https://docs.docker.com/language/nodejs/develop/
# CMD ["ng", "serve", "--host", "0.0.0.0"]
