# Copyright 2020 Coinbase, Inc.
# Copyright 2021 Rosetta Dogecoin Developers
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Build dogecoind
FROM ubuntu:18.04 as dogecoind-builder

RUN mkdir -p /app \
  && chown -R nobody:nogroup /app
WORKDIR /app

RUN apt-get update && apt-get install -y curl
ENV DOGECOIN_VERSION 1.14.3
ENV DOGECOIN_DOWNLOAD_SHA256 a95cc29ac3c19a450e9083cc3ac24b6f61763d3ed1563bfc3ea9afbf0a2804fd
ENV DOGECOIN_DOWNLOAD_URL https://github.com/dogecoin/dogecoin/releases/download/v$DOGECOIN_VERSION/dogecoin-$DOGECOIN_VERSION-x86_64-linux-gnu.tar.gz

# Fetch and verify source
RUN curl -fsSL "$DOGECOIN_DOWNLOAD_URL" -o dogecoin.tar.gz \
  && echo "$DOGECOIN_DOWNLOAD_SHA256 dogecoin.tar.gz" | sha256sum -c \
  && tar -xzf dogecoin.tar.gz dogecoin-$DOGECOIN_VERSION/bin/dogecoind \
  && rm dogecoin.tar.gz \
  && mv dogecoin-$DOGECOIN_VERSION/bin/dogecoind /app/dogecoind \
  && rm -rf dogecoin-$DOGECOIN_VERSION

# Build Rosetta Server Components
FROM ubuntu:18.04 as rosetta-builder

RUN mkdir -p /app \
  && chown -R nobody:nogroup /app
WORKDIR /app

RUN apt-get update && apt-get install -y curl make gcc g++
ENV GOLANG_VERSION 1.15.5
ENV GOLANG_DOWNLOAD_SHA256 9a58494e8da722c3aef248c9227b0e9c528c7318309827780f16220998180a0d
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
  && echo "$GOLANG_DOWNLOAD_SHA256 golang.tar.gz" | sha256sum -c - \
  && tar -C /usr/local -xzf golang.tar.gz \
  && rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Get dependencies first
COPY go.mod go.sum src/
RUN cd src && go mod download

# Use native remote build context to build in any directory
COPY . src
RUN cd src \
  && go build \
  && cd .. \
  && mv src/rosetta-dogecoin /app/rosetta-dogecoin \
  && mv src/assets/* /app \
  && rm -rf src

## Build Final Image
FROM ubuntu:18.04

RUN mkdir -p /app \
  && chown -R nobody:nogroup /app \
  && mkdir -p /data \
  && chown -R nobody:nogroup /data

WORKDIR /app

# Copy binary from dogecoind-builder
COPY --from=dogecoind-builder /app/dogecoind /app/dogecoind

# Copy binary from rosetta-builder
COPY --from=rosetta-builder /app/* /app/

# Set permissions for everything added to /app
RUN chmod -R 755 /app/*

CMD ["/app/rosetta-dogecoin"]
