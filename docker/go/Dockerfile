FROM golang:1.22.5 AS base_app

WORKDIR /app 

RUN apt-get update && apt-get install -y \
  make

ARG user=appuser
ARG group=appuser
ARG uid=1000
ARG gid=1000
RUN groupadd -g ${gid} ${group}
RUN useradd -u ${uid} -g ${group} -s /bin/sh -m ${user}
# Switch to user
USER ${uid}:${gid}

RUN go install github.com/google/wire/cmd/wire@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.0
RUN go install github.com/rubenv/sql-migrate/...@latest

COPY go.mod go.sum /app/

RUN go mod download

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
