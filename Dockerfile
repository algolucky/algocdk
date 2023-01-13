FROM docker.io/library/golang:1.19.5-bullseye

# hadolint ignore=DL4006,DL3008
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    curl -sL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y --no-install-recommends nodejs && \
    rm -rf /var/lib/apt/lists/* && \
    node -v && \
    npm -v && \
    npm install -g cdktf-cli@0.15.2 && \
    cdktf --version

COPY . $GOPATH/src/algocdk

WORKDIR $GOPATH/src/algocdk

RUN make install

WORKDIR $GOPATH

ENTRYPOINT [ "algocdk" ]
