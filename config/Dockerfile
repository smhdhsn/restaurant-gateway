# Building binary.
FROM golang:1.18 AS binary

RUN apt-get update && apt-get upgrade -yq

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN make build_all BIN_DIR=/build

# Building container to run binary.
FROM debian:bullseye

COPY --from=binary /build/config/ /app/config/
COPY --from=binary /build/server /app/

WORKDIR /app

EXPOSE 9000
